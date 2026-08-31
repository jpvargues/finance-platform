import { useState, useEffect } from "react";
import { useParams, Link } from "react-router-dom";

interface ETF {
  id: number;
  ticker: string;
  name: string;
  isin: string;
  is_accumulating: boolean;
  category: string;
}

function EtfDetail() {
  const { id } = useParams();
  const [etf, setEtf] = useState<ETF | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    fetch(`http://localhost:8080/etfs/${id}`)
      .then((response) => {
        if (!response.ok) throw new Error("ETF not found");
        return response.json();
      })
      .then((data) => setEtf(data))
      .catch(() => setError("Error loading ETF"))
      .finally(() => setLoading(false));
  }, [id]);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>{error}</p>;
  if (!etf) return null;

  return (
    <>
      <Link to="/">← Back to list</Link>
      <h1>{etf.ticker}</h1>
      <p>{etf.name}</p>
      <p>ISIN: {etf.isin}</p>
      <p>Category: {etf.category}</p>
      <p>{etf.is_accumulating ? "Accumulating" : "Distributing"}</p>
    </>
  );
}

export default EtfDetail;
