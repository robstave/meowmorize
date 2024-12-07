import React, { useEffect, useRef } from 'react';
import * as d3 from 'd3';

const HorizontalProgressBar = ({ index, total }) => {
  const svgRef = useRef();
  const tooltipRef = useRef();

  useEffect(() => {
    const currentIndex = Number(index) || 0;
    const totalCount = Number(total) || 0;

    d3.select(svgRef.current).selectAll("*").remove();

    const width = 300;
    const height = 40;
    const margin = { top: 10, right: 10, bottom: 10, left: 0 };
    const cornerRadius = 6; // Added corner radius

    const progress = totalCount > 0 ? currentIndex / totalCount : 0;
    
    const data = [
      { value: progress, color: '#2e7d32', label: 'Progress', count: currentIndex },
      { value: 1 - progress, color: '#f5f5f5', label: 'Remaining', count: totalCount - currentIndex }
    ];

    const svg = d3.select(svgRef.current)
      .attr('width', width)
      .attr('height', height);

    const xScale = d3.scaleLinear()
      .domain([0, 1])
      .range([margin.left, width - margin.right]);

    let cumulative = 0;
    data.forEach(d => {
      d.start = cumulative;
      cumulative += d.value;
      d.end = cumulative;
    });

    if (!tooltipRef.current) {
      tooltipRef.current = d3.select('body')
        .append('div')
        .style('position', 'absolute')
        .style('pointer-events', 'none')
        .style('background-color', '#333')
        .style('color', '#fff')
        .style('padding', '6px 8px')
        .style('border-radius', '4px')
        .style('font-size', '12px')
        .style('z-index', '1000')
        .style('white-space', 'nowrap')
        .style('display', 'none');
    }

    const tooltip = tooltipRef.current;

    svg.selectAll('rect')
      .data(data)
      .enter()
      .append('rect')
      .attr('x', d => xScale(d.start))
      .attr('y', margin.top)
      .attr('width', d => xScale(d.end) - xScale(d.start))
      .attr('height', height - margin.top - margin.bottom)
      .attr('rx', cornerRadius) // Added corner radius
      .attr('ry', cornerRadius) // Added corner radius
      .attr('fill', d => d.color)
      .on('mouseover', (event, d) => {
        tooltip
          .style('left', `${event.clientX + 10}px`)
          .style('top', `${event.clientY - 28}px`)
          .html(`${d.label}: ${d.count}`)
          .style('display', 'block');
      })
      .on('mousemove', (event) => {
        tooltip
          .style('left', `${event.clientX + 10}px`)
          .style('top', `${event.clientY - 28}px`);
      })
      .on('mouseout', () => {
        tooltip.style('display', 'none');
      });

    return () => {
      if (tooltipRef.current) {
        tooltipRef.current.remove();
        tooltipRef.current = null;
      }
    };
  }, [index, total]);

  return (
    <div className="relative">
      <svg ref={svgRef}></svg>
    </div>
  );
};

export default HorizontalProgressBar;