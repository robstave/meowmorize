import React, { useState } from 'react';

const CatStatusChart = ({ pass = 0, skip = 0, fail = 0, retire = false }) => {
  const [tooltip, setTooltip] = useState({ show: false, content: '', x: 0, y: 0 });

  const width = 80;
  const height = 80;
  const center = { x: width / 2, y: height / 2 };
  const radius = 30; // Radius for the pie chart

  // Calculate total and prepare data
  const total = pass + skip + fail || 1;
  const data = retire 
    ? [{ label: 'Retired', value: 1, color: '#808080' }]
    : [
        { label: 'Pass', value: pass, color: '#2e7d32' },
        { label: 'Skip', value: skip, color: '#ed6c02' },
        { label: 'Fail', value: fail, color: '#d32f2f' }
      ].filter(item => item.value > 0);

  const finalData = data.length > 0 ? data : [{ label: 'No Data', value: 1, color: '#ccc' }];

  // Calculate segments angles
  let startAngle = -Math.PI / 2;
  const segments = finalData.map(item => {
    const angle = (item.value / total) * 2 * Math.PI;
    const segment = {
      ...item,
      startAngle,
      endAngle: startAngle + angle,
    };
    startAngle += angle;
    return segment;
  });

  // Create pie segment paths
  const createPieSegment = (segment) => {
    const startX = center.x + radius * Math.cos(segment.startAngle);
    const startY = center.y + radius * Math.sin(segment.startAngle);
    const endX = center.x + radius * Math.cos(segment.endAngle);
    const endY = center.y + radius * Math.sin(segment.endAngle);

    const largeArcFlag = segment.endAngle - segment.startAngle > Math.PI ? 1 : 0;

    return `
      M ${center.x} ${center.y}
      L ${startX} ${startY}
      A ${radius} ${radius} 0 ${largeArcFlag} 1 ${endX} ${endY}
      Z
    `;
  };

  return (
    <div style={{ position: 'relative', width, height }}>
      <svg width={width} height={height}>
        {/* 1. Static Ears (Drawn First) */}
        {/* Left Ear */}
        <path
          d={`
            M ${center.x - 25} ${center.y - 10}
            L ${center.x - 25} ${center.y - 40}
            L ${center.x - 5} ${center.y - 30}
            Z
          `}
          fill="#000" // Static black color
        />

        {/* Right Ear */}
        <path
          d={`
            M ${center.x + 25} ${center.y - 10}
            L ${center.x + 25} ${center.y - 40}
            L ${center.x + 5} ${center.y - 30}
            Z
          `}
          fill="#000" // Static black color
        />

        {/* 2. Main Pie Segments or Full Circle */}
        {segments.length === 1 ? (
          // Single Segment: Render a full circle
          <circle
            cx={center.x}
            cy={center.y}
            r={radius}
            fill={segments[0].color}
            onMouseMove={(e) => {
              const rect = e.currentTarget.getBoundingClientRect();
              setTooltip({
                show: true,
                content: `${segments[0].label}: ${segments[0].value}`,
                x: e.clientX - rect.left + 10,
                y: e.clientY - rect.top - 20
              });
            }}
            onMouseLeave={() => setTooltip({ ...tooltip, show: false })}
          />
        ) : (
          // Multiple Segments: Render each segment as a path
          segments.map((segment, i) => (
            <path
              key={i}
              d={createPieSegment(segment)}
              fill={segment.color}
              onMouseMove={(e) => {
                const rect = e.currentTarget.getBoundingClientRect();
                setTooltip({
                  show: true,
                  content: `${segment.label}: ${segment.value}`,
                  x: e.clientX - rect.left + 10,
                  y: e.clientY - rect.top - 20
                });
              }}
              onMouseLeave={() => setTooltip({ ...tooltip, show: false })}
            />
          ))
        )}

        {/* 3. Thick Black Border Around Pie Chart (Drawn Above Pie) */}
        <circle
          cx={center.x}
          cy={center.y}
          r={radius}
          fill="none"
          stroke="#000"
          strokeWidth="2" // Adjust thickness as needed
        />

        {/* 4. Cat Facial Features (Drawn on Top) */}
        {/* Eyes */}
        <circle cx={center.x - 12} cy={center.y - 5} r={3} fill="#333" /> {/* Left eye */}
        <circle cx={center.x + 12} cy={center.y - 5} r={3} fill="#333" /> {/* Right eye */}
        {/* Nose */}
        <circle cx={center.x} cy={center.y + 5} r={2} fill="#333" /> {/* Nose */}
        
        {/* Mouth */}
        <path 
          d={`M ${center.x - 8} ${center.y + 10} Q ${center.x} ${center.y + 15} ${center.x + 8} ${center.y + 10}`}
          stroke="#333"
          fill="none"
          strokeWidth="2"
        />
        
        {/* Whiskers */}
        <g stroke="#333" strokeWidth="1.5">
          <line x1={center.x - 20} y1={center.y + 5} x2={center.x - 10} y2={center.y + 8} />
          <line x1={center.x - 20} y1={center.y + 10} x2={center.x - 10} y2={center.y + 10} />
          <line x1={center.x + 20} y1={center.y + 5} x2={center.x + 10} y2={center.y + 8} />
          <line x1={center.x + 20} y1={center.y + 10} x2={center.x + 10} y2={center.y + 10} />
        </g>
      </svg>
      
      {/* Tooltip */}
      {tooltip.show && (
        <div
          style={{
            position: 'absolute',
            left: tooltip.x,
            top: tooltip.y,
            backgroundColor: '#333',
            color: '#fff',
            padding: '4px 8px',
            borderRadius: '4px',
            fontSize: '12px',
            pointerEvents: 'none',
            whiteSpace: 'nowrap',
            zIndex: 1000
          }}
        >
          {tooltip.content}
        </div>
      )}
    </div>
  );
};

export default CatStatusChart;
