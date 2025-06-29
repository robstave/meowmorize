import React, { useEffect, useRef } from 'react';
import * as d3 from 'd3';
import './CardStatsBar.css';

const CardStatsBar = ({ cards = []}) => {
  const chartRef = useRef();

  useEffect(() => {
    // Clear any existing content
    d3.select(chartRef.current).selectAll("*").remove();

    // Dimensions and configuration
    const totalCards = cards.length;
    const svgWidth = 500;
    const svgHeight = 120;
    const barHeight = 16;
    const starBlockHeight = 4;
    const starGap = 2;          // Gap between main bar and star bars
    const baseY = svgHeight - barHeight - 10;

    // Create SVG
    const svg = d3.select(chartRef.current)
      .append("svg")
      .attr("width", svgWidth)
      .attr("height", svgHeight);

    // Tooltip
    const tooltip = d3.select("body")
      .append("div")
      .attr("class", "card-stats-tooltip")
      .style("opacity", 0);

    // Color function for the main bar
    function getColor(card) {
      if (!card.viewed) return '#808080';
      if (card.passed)  return '#2e7d32';
      if (card.failed)  return '#d32f2f';
      if (card.skipped) return '#ed6c02';
      return '#808080';
    }

    // Aggregated totals for tooltip (optional)
    const totalViewed  = d3.sum(cards, d => d.viewed ? 1 : 0);
    const totalPassed  = d3.sum(cards, d => d.passed ? 1 : 0);
    const totalFailed  = d3.sum(cards, d => d.failed ? 1 : 0);
    const totalSkipped = d3.sum(cards, d => d.skipped ? 1 : 0);

    // X-scale
    const xScale = d3.scaleLinear()
      .domain([0, totalCards])
      .range([0, svgWidth]);

    const cardWidth = svgWidth / totalCards;

    // Create a group for each card
    const cardGroups = svg.selectAll(".card-group")
      .data(cards)
      .enter()
      .append("g")
      .attr("class", "card-group")
      .attr("transform", (d, i) => `translate(${xScale(i)},0)`)
      .on("mouseover", function(event, d) {
        const tooltipContent = `
          <strong>Card ID:</strong> ${d.card_id}<br>
          <strong>Stars:</strong> ${d.stars}<br>
          <hr style="border:none;border-top:1px solid #aaa;margin:4px 0;">
          <strong>Total Viewed:</strong> ${totalViewed}<br>
          <strong>Total Passed:</strong> ${totalPassed}<br>
          <strong>Total Failed:</strong> ${totalFailed}<br>
          <strong>Total Skipped:</strong> ${totalSkipped}
        `;
        tooltip.html(tooltipContent)
          .style("left", `${event.pageX + 10}px`)
          .style("top", `${event.pageY - 10}px`)
          .style("opacity", 1);
      })
      .on("mouseout", function() {
        tooltip.style("opacity", 0);
      });

    // Main progress bar for each card
    cardGroups.append("rect")
      .attr("class", "main-bar")
      .attr("x", 0)
      .attr("y", baseY)
      .attr("width", cardWidth)
      .attr("height", barHeight)
      .attr("fill", d => getColor(d));

    // Stacked star bars for each card
    cardGroups.each(function(d) {
      const g = d3.select(this);
      for (let s = 0; s < d.stars; s++) {
        g.append("rect")
          .attr("class", "star-bar")
          .attr("x", 0)
          .attr("y", baseY - starGap - (s + 1) * starBlockHeight)
          .attr("width", cardWidth)
          .attr("height", starBlockHeight)
          .attr("fill", "gold");
      }
    });

    // Cleanup tooltip on unmount
    return () => {
      tooltip.remove();
    };
  }, [cards]);

  return <div ref={chartRef}></div>;
};

export default CardStatsBar;
