import React, { useEffect, useRef } from 'react';
import * as d3 from 'd3';

const CardStatsBar = ({ cards }) => {
  const svgRef = useRef();

  useEffect(() => {
    // Clear previous render
    d3.select(svgRef.current).selectAll("*").remove();
    
    const width = 300;  // Adjust as needed
    const height = 16;
    const total = cards.length;

    if (total === 0) {
      return; // No cards, no need to render anything
    }

    // Define a function to determine card color
    const getColor = (card) => {
      if (!card.viewed) return '#808080';    // Grey
      if (card.passed) return '#2e7d32';     // Green
      if (card.failed) return '#d32f2f';     // Red
      if (card.skipped) return '#ed6c02';    // Yellow
      // If viewed but none of the above states:
      return '#808080';
    };

   

    // Create SVG
    const svg = d3.select(svgRef.current)
      .attr('width', width)
      .attr('height', height);

    // Create tooltip
    const body = d3.select('body');
    const tooltip = body.append('div')
      .attr('class', 'absolute hidden bg-gray-800 text-white p-2 rounded text-xs')
      .style('pointer-events', 'none');

    // Scales
    const xScale = d3.scaleLinear()
      .domain([0, total])
      .range([0, width]);

    // Container for the cards bar
    const barGroup = svg.append('g').attr('class', 'cards-bar');

    // Draw rectangles for each card
    barGroup.selectAll('rect')
      .data(cards)
      .enter()
      .append('rect')
      .attr('x', (d, i) => xScale(i))
      .attr('y', 0)
      .attr('width', width / total)
      .attr('height', height)
      .attr('fill', d => getColor(d))
      .on('mouseover', (event, d) => {
        const tooltipContent = `
          CardID: ${d.CardID}<br/>
          Viewed: ${d.Viewed}<br/>
          Passed: ${d.Passed}<br/>
          Failed: ${d.Failed}<br/>
          Skipped: ${d.Skipped}
        `;
        
        tooltip
          .style('left', `${event.pageX + 10}px`)
          .style('top', `${event.pageY - 10}px`)
          .html(tooltipContent)
          .classed('hidden', false);
      })
      .on('mouseout', () => {
        tooltip.classed('hidden', true);
      });

    // Cleanup on unmount
    return () => {
      tooltip.remove();
    };
  }, [cards]);

  return (
    <div className="inline-block">
      <svg ref={svgRef}></svg>
    </div>
  );
};

export default CardStatsBar;
