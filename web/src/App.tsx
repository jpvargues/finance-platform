import { useState, useEffect } from "react";
import "./App.css";

interface ETF {
  id: number;
  ticker: string;
  name: string;
  isin: string;
  is_accumulating: boolean;
  category: string;
}

function App() {
  const [etfs, setEtfs] = useState<ETF[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/etfs")
      .then((response) => response.json())
      .then((data) => setEtfs(data));
  }, []);

  return (
    <>
      <h1>ETF Explorer</h1>
      <ul>
        {etfs.map((etf) => (
          <li key={etf.id}>
            {etf.ticker} — {etf.name} ({etf.category})
          </li>
        ))}
      </ul>
    </>
  );
}

export default App;
