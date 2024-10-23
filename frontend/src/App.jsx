
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css'

function App() {
  const [ matches , setMatches ] = useState([])


  useEffect(() => {
      fetchMatches()

      const ws = new WebSocket("ws://localhost:8081/ws")

      ws.onopen = function(event) {
        console.log('WebSocket connection established.');
        
        // Example: send a message
        ws.send('Hello, WebSocket Server!');
      };

      ws.onmessage = (event) => {
        console.log(event)
        const updatedMatch = JSON.parse(event.data)
        setMatches((prev) => (
          prev.map(match => (
            match.id === updatedMatch.id  ? updatedMatch : match
          ))
        ))
      };

      ws.onclose = (event) => {
        console.log("websocket connection closed")
      }

      ws.onerror = (error) => {
        console.log(error)
      }      

      // return () => {
      //   ws.close()
      // }

  },[]);

  const fetchMatches = async () => {
    try {
      const matches = await axios.get("http://127.0.0.1:4000/match/")
      console.log(matches)
      setMatches(matches.data.data)      
    } catch (error) {
        console.log('Error Fetching matches: ', error)
    }
  }

  return (
    <div className='app'>
        <div className="live-scoreboard">
            <h1>Live Scoreboard</h1>
            {matches.map(match => (
                <div key={match.id} className="match">
                    <h2>{match.home_team} vs {match.away_team}</h2>
                    <p>Score: {match.home_score} - {match.away_score}</p>
                    <p>Status: {match.status}</p>
                </div>
            ))}
        </div>
    </div>
  )
}

export default App
