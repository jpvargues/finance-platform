import { BrowserRouter, Routes, Route } from "react-router-dom";
import EtfList from "./EtfList";
import EtfDetail from "./EtfDetail";
import "./App.css";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<EtfList />} />
        <Route path="/etfs/:id" element={<EtfDetail />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
