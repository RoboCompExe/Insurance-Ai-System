
import { useState } from "react";
import axios from "axios";

export default function Home() {
  const [data, setData] = useState(null);

  const run = async () => {
    const res = await axios.post("http://localhost:8080/fraud-check");
    setData(res.data);
  };

  return (
    <div style={{padding:20}}>
      <h1>AI Fraud Dashboard</h1>
      <button onClick={run}>Check</button>
      {data && <pre>{JSON.stringify(data,null,2)}</pre>}
    </div>
  );
}
