import React, { useCallback, useEffect, useMemo, useRef, useState } from 'react';
import { Box, Button, Paper, Stack, Typography } from '@mui/material';

const PANEL_SIZE = 260;
const PLAYER_X = PANEL_SIZE / 2;
const PLAYER_Y = PANEL_SIZE / 2;
const SPHERE_RADIUS = 14;
const ORBIT_RADIUS = 90;
const SAMPLE_RATE = 8000;

const byteBeatSample = (t) => {
  const value = t * (t ^ t + (((t >> 15) | 1) ^ ((t - 1280) ^ t) >> 10));
  return ((value & 255) / 127.5) - 1;
};

const ByteBeatSphere = () => {
  const audioContextRef = useRef(null);
  const sourceNodeRef = useRef(null);
  const gainNodeRef = useRef(null);
  const [audioEnabled, setAudioEnabled] = useState(false);
  const [angle, setAngle] = useState(0);
  const [playerPosition, setPlayerPosition] = useState({ x: PLAYER_X, y: PLAYER_Y });

  const spherePosition = useMemo(() => ({
    x: PLAYER_X + ORBIT_RADIUS * Math.cos(angle),
    y: PLAYER_Y + ORBIT_RADIUS * Math.sin(angle),
  }), [angle]);

  const distance = useMemo(() => {
    const dx = spherePosition.x - playerPosition.x;
    const dy = spherePosition.y - playerPosition.y;
    return Math.sqrt((dx * dx) + (dy * dy));
  }, [spherePosition, playerPosition]);

  const normalizedVolume = useMemo(() => {
    const maxDistance = Math.sqrt((PANEL_SIZE * PANEL_SIZE) + (PANEL_SIZE * PANEL_SIZE));
    return Math.max(0, Math.min(1, 1 - (distance / maxDistance)));
  }, [distance]);

  const waveform = useMemo(() => {
    const points = [];
    const width = 360;
    const height = 90;

    for (let i = 0; i < width; i += 1) {
      const t = Math.floor((i / width) * SAMPLE_RATE * 0.06);
      const y = ((byteBeatSample(t) + 1) / 2) * height;
      points.push(`${i},${height - y}`);
    }

    return points.join(' ');
  }, []);

  const ensureAudio = useCallback(async () => {
    if (audioContextRef.current) {
      if (audioContextRef.current.state === 'suspended') {
        await audioContextRef.current.resume();
      }
      return;
    }

    const context = new window.AudioContext();
    const gainNode = context.createGain();
    gainNode.gain.value = 0;

    const frameCount = context.sampleRate;
    const audioBuffer = context.createBuffer(1, frameCount, context.sampleRate);
    const channelData = audioBuffer.getChannelData(0);

    for (let i = 0; i < frameCount; i += 1) {
      const t = Math.floor((i / context.sampleRate) * SAMPLE_RATE);
      channelData[i] = byteBeatSample(t);
    }

    const source = context.createBufferSource();
    source.buffer = audioBuffer;
    source.loop = true;

    source.connect(gainNode);
    gainNode.connect(context.destination);
    source.start();

    audioContextRef.current = context;
    sourceNodeRef.current = source;
    gainNodeRef.current = gainNode;
  }, []);

  useEffect(() => {
    let animationFrameId;
    let start;

    const step = (timestamp) => {
      if (!start) {
        start = timestamp;
      }

      const elapsed = (timestamp - start) / 1000;
      setAngle(elapsed * 1.2);
      animationFrameId = window.requestAnimationFrame(step);
    };

    animationFrameId = window.requestAnimationFrame(step);

    return () => {
      if (animationFrameId) {
        window.cancelAnimationFrame(animationFrameId);
      }
    };
  }, []);

  useEffect(() => {
    if (!gainNodeRef.current) {
      return;
    }

    gainNodeRef.current.gain.linearRampToValueAtTime(
      normalizedVolume * 0.35,
      audioContextRef.current.currentTime + 0.08,
    );
  }, [normalizedVolume]);

  useEffect(() => () => {
    sourceNodeRef.current?.stop();
    sourceNodeRef.current?.disconnect();
    gainNodeRef.current?.disconnect();
    audioContextRef.current?.close();
  }, []);

  const handleEnableAudio = async () => {
    await ensureAudio();
    setAudioEnabled(true);
  };

  const handleMouseMove = (event) => {
    const rect = event.currentTarget.getBoundingClientRect();
    const x = event.clientX - rect.left;
    const y = event.clientY - rect.top;
    setPlayerPosition({
      x: Math.max(0, Math.min(PANEL_SIZE, x)),
      y: Math.max(0, Math.min(PANEL_SIZE, y)),
    });
  };

  return (
    <Paper sx={{ p: 2, mb: 2 }}>
      <Stack spacing={1.5}>
        <Typography variant="h6">ByteBeat Sphere Prototype</Typography>
        <Typography variant="body2" color="text.secondary">
          Move your cursor (player) and listen to the sphere enemy get louder as it gets closer.
        </Typography>

        <Box
          onMouseMove={handleMouseMove}
          sx={{
            position: 'relative',
            width: PANEL_SIZE,
            height: PANEL_SIZE,
            borderRadius: 2,
            border: '1px solid',
            borderColor: 'divider',
            background: 'radial-gradient(circle at center, rgba(120,120,120,0.25), rgba(20,20,20,0.85))',
            overflow: 'hidden',
          }}
        >
          <Box
            sx={{
              position: 'absolute',
              left: playerPosition.x - 8,
              top: playerPosition.y - 8,
              width: 16,
              height: 16,
              borderRadius: '50%',
              backgroundColor: 'primary.main',
              boxShadow: '0 0 8px rgba(50,130,255,0.9)',
            }}
          />
          <Box
            sx={{
              position: 'absolute',
              left: spherePosition.x - SPHERE_RADIUS,
              top: spherePosition.y - SPHERE_RADIUS,
              width: SPHERE_RADIUS * 2,
              height: SPHERE_RADIUS * 2,
              borderRadius: '50%',
              background: 'radial-gradient(circle at 30% 30%, #ffb3d1, #d81b60)',
              boxShadow: `0 0 ${6 + (normalizedVolume * 20)}px rgba(216,27,96,0.95)`,
            }}
          />
        </Box>

        <svg viewBox="0 0 360 90" style={{ width: '100%', maxWidth: 360, border: '1px solid #555', borderRadius: 8 }}>
          <polyline fill="none" stroke="#3f51b5" strokeWidth="2" points={waveform} />
        </svg>

        <Typography variant="caption" color="text.secondary">
          Formula: <code>t*(t^t+(t&gt;&gt;15|1)^(t-1280^t)&gt;&gt;10)</code> @ 8 kHz
        </Typography>

        <Typography variant="body2">
          Distance: {distance.toFixed(1)} px · Volume: {(normalizedVolume * 100).toFixed(0)}%
        </Typography>

        <Button variant="contained" onClick={handleEnableAudio}>
          {audioEnabled ? 'Audio Active' : 'Enable ByteBeat Audio'}
        </Button>
      </Stack>
    </Paper>
  );
};

export default ByteBeatSphere;
