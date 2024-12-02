import React, { useEffect, useRef } from 'react';
import * as d3 from 'd3';

const HorizontalStatusBar = ({ pass = 10, skip = 4, fail = 6, retire = false }) => {
  const svgRef = useRef();
  
  useEffect(() => {
    // Clear any existing content
    d3.select(svgRef.current).selectAll("*").remove();
    
    // Set up dimensions
    const width = 400;
    const height = 60;
    const margin = { top: 20, right: 20, bottom: 20, left: 20 };
    
    // Calculate total for percentages
    const total = pass + skip + fail;
    
    // Create SVG
    const svg = d3.select(svgRef.current)
      .attr('width', width)
      .attr('height', height);
    
    // Create data structure
    const data = retire ? 
      [{ value: 1, color: '#808080', label: 'Retired', count: total }] :
      [
        { value: pass/total, color: '#4CAF50', label: 'Pass', count: pass },
        { value: skip/total, color: '#FFC107', label: 'Skip', count: skip },
        { value: fail/total, color: '#F44336', label: 'Fail', count: fail }
      ];
    
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
    
    // Create tooltip
    const tooltip = d3.select('body')
      .append('div')
      .attr('class', 'absolute hidden bg-gray-800 text-white p-2 rounded text-sm')
      .style('pointer-events', 'none');
    
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
          .style('left', `${event.pageX + 10}px`)
          .style('top', `${event.pageY - 10}px`)
          .html(`${d.label}: ${d.count}`)
          .classed('hidden', false);
      })
      .on('mouseout', () => {
        tooltip.classed('hidden', true);
      });
    
    // Cleanup function
    return () => {
      tooltip.remove();
    };
  }, [pass, skip, fail, retire]); // Dependencies array
  
  return (
    <div className="relative">
      <svg ref={svgRef}></svg>
    </div>
  );
};

export default HorizontalStatusBar;