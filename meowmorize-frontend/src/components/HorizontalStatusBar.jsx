// src/components/HorizontalStatusBar.jsx

// src/components/HorizontalStatusBar.jsx
import React, { useEffect, useRef } from 'react';
import * as d3 from 'd3';
import PropTypes from 'prop-types';

const HorizontalStatusBar = ({ pass, skip, fail, retire = false }) => {
  const svgRef = useRef();
  const tooltipRef = useRef(); // Reference to the tooltip

  useEffect(() => {
    // Ensure pass, skip, fail are numbers
    const passNum = Number(pass) || 0;
    const skipNum = Number(skip) || 0;
    const failNum = Number(fail) || 0;

    // Clear any existing SVG content
    d3.select(svgRef.current).selectAll("*").remove();

    // Set up dimensions
    const width = 300;
    const height = 40;
    const margin = { top: 10, right: 10, bottom: 10, left: 0 };

    // Calculate total for percentages
    const total = passNum + skipNum + failNum;

    // Handle retire case
    const data = retire
      ? [{ value: 1, color: '#808080', label: 'Retired', count: total }]
      : [
        { value: passNum / total, color: '#2e7d32', label: 'Pass', count: passNum },
        { value: skipNum / total, color: '#ed6c02 ', label: 'Skip', count: skipNum },
        { value: failNum / total, color: '#d32f2f', label: 'Fail', count: failNum }
        ];

    // Create SVG
    const svg = d3.select(svgRef.current)
      .attr('width', width)
      .attr('height', height);

    // Create scales
    const xScale = d3.scaleLinear()
      .domain([0, 1])
      .range([margin.left, width - margin.right]);

    // Calculate cumulative values for x positions
    let cumulative = 0;
    data.forEach(d => {
      d.start = cumulative;
      cumulative += d.value;
      d.end = cumulative;
    });

    // Create Tooltip if it doesn't exist
    if (!tooltipRef.current) {
      tooltipRef.current = d3.select('body')
        .append('div')
        .attr('class', 'tooltip') // Optional: for additional styling
        .style('position', 'absolute')
        .style('pointer-events', 'none')
        .style('background-color', '#333')
        .style('color', '#fff')
        .style('padding', '6px 8px')
        .style('border-radius', '4px')
        .style('font-size', '12px')
        .style('z-index', '1000')
        .style('white-space', 'nowrap')
        .style('display', 'none'); // Hidden by default
    }

    const tooltip = tooltipRef.current;

    // Draw bars
    svg.selectAll('rect')
      .data(data)
      .enter()
      .append('rect')
      .attr('x', d => xScale(d.start))
      .attr('y', margin.top)
      .attr('width', d => xScale(d.end) - xScale(d.start))
      .attr('height', height - margin.top - margin.bottom)
      .attr('fill', d => d.color)
      .on('mouseover', (event, d) => {
        tooltip
          .style('left', `${event.clientX + 10}px`)
          .style('top', `${event.clientY - 28}px`)
          .html(`${d.label}: ${d.count}`)
          .style('display', 'block'); // Show tooltip
      })
      .on('mousemove', (event) => {
        tooltip
          .style('left', `${event.clientX + 10}px`)
          .style('top', `${event.clientY - 28}px`);
      })
      .on('mouseout', () => {
        tooltip.style('display', 'none'); // Hide tooltip
      });

    // Cleanup function to remove tooltip on unmount
    return () => {
      if (tooltipRef.current) {
        tooltipRef.current.remove();
        tooltipRef.current = null;
      }
    };
  }, [pass, skip, fail, retire]); // Dependencies array

  return (
    <div className="relative">
      <svg ref={svgRef}></svg>
    </div>
  );
};

// Define prop types to enforce required props
HorizontalStatusBar.propTypes = {
  pass: PropTypes.number.isRequired,
  skip: PropTypes.number.isRequired,
  fail: PropTypes.number.isRequired,
  retire: PropTypes.bool,
};

export default HorizontalStatusBar;

 