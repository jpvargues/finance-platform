import { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import "./App.css";

interface ETF {
  id: number;
  ticker: string;
  name: string;
  isin: string;
  is_accumulating: boolean;
  category: string;
}

function EtfList() {
  const [etfs, setEtfs] = useState<ETF[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    fetch("http://localhost:8080/etfs")
      .then((response) => response.json())
      .then((data) => setEtfs(data))
      .catch(() => setError("Error loading ETFs"))
      .finally(() => setLoading(false));
  }, []);

  return (
    <>
      <h1>ETF Explorer</h1>
      {loading && <p>Loading...</p>}
      {error && <p>{error}</p>}
      <ul>
        {etfs.map((etf) => (
          <li key={etf.id}>
            <Link to={`/etfs/${etf.id}`}>
              {etf.ticker} — {etf.name} ({etf.category})
            </Link>
          </li>
        ))}
      </ul>
    </>
  );
}

export default EtfList;
