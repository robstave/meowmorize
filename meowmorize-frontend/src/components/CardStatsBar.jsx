import React, { useEffect, useRef } from 'react';
import * as d3 from 'd3';
import './CardStatsBar.css';

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

    // Calculate total counts
    const totalViewed = d3.sum(cards, d => d.viewed ? 1 : 0);
    const totalPassed = d3.sum(cards, d => d.passed ? 1 : 0);
    const totalFailed = d3.sum(cards, d => d.failed ? 1 : 0);
    const totalSkipped = d3.sum(cards, d => d.skipped ? 1 : 0);

    // Create SVG
    const svg = d3.select(svgRef.current)
      .attr('width', width)
      .attr('height', height);

    // Create tooltip
    const tooltip = d3.select('body').append('div')
      .attr('class', 'card-stats-tooltip')
      .style('opacity', 0);

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
          <br/>
          <strong>Total Viewed:</strong> ${totalViewed}<br/>
          <strong>Total Passed:</strong> ${totalPassed}<br/>
          <strong>Total Failed:</strong> ${totalFailed}<br/>
          <strong>Total Skipped:</strong> ${totalSkipped}
        `;
        
        tooltip
          .html(tooltipContent)
          .style('left', `${event.pageX + 10}px`)
          .style('top', `${event.pageY - 10}px`)
          .style('opacity', 1);
      })
      .on('mouseout', () => {
        tooltip.style('opacity', 0);
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
