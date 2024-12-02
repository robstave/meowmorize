import React, { useEffect, useRef } from 'react';
import * as d3 from 'd3';

const StatusSparkline = ({ pass, skip, fail, retire = false }) => {
  const svgRef = useRef();
  
  useEffect(() => {
    // Clear any existing content
    d3.select(svgRef.current).selectAll("*").remove();
    
    // Set up dimensions
    const width = 120;
    const height = 16;
    
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
      .range([0, width]);
    
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
      .attr('class', 'absolute hidden bg-gray-800 text-white p-2 rounded text-xs')
      .style('pointer-events', 'none');
    
    // Create container group for the sparkline
    const sparkline = svg
      .append('g')
      .attr('class', 'sparkline');
    
    // Draw bars
    sparkline.selectAll('rect')
      .data(data)
      .enter()
      .append('rect')
      .attr('x', d => xScale(d.start))
      .attr('y', 0)
      .attr('width', d => xScale(d.end) - xScale(d.start))
      .attr('height', height)
      .attr('fill', d => d.color);
    
    // Add mouse interaction for the entire sparkline
    svg.append('rect')
      .attr('width', width)
      .attr('height', height)
      .attr('fill', 'transparent')
      .on('mouseover', (event) => {
        const tooltipContent = retire ?
          `Retired: ${total}` :
          `Pass: ${pass} | Skip: ${skip} | Fail: ${fail} | Total: ${total}`;
        
        tooltip
          .style('left', `${event.pageX + 10}px`)
          .style('top', `${event.pageY - 10}px`)
          .html(tooltipContent)
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
    <div className="inline-block">
      <svg ref={svgRef}></svg>
    </div>
  );
};

export default StatusSparkline;