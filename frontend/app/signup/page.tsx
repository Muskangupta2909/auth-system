"use client";
import { useState } from "react";

export default function Signup() {
  const [email, setEmail] = useState("");
  const [mobile, setMobile] = useState("");
  const [password, setPassword] = useState("");

  const submit = async () => {
    const res = await fetch("http://localhost:8080/signup", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, mobile, password }),
    });
    const data = await res.json();
    alert(data.message);
  };

  return (
    <div style={{ padding: 40 }}>
      <h2>Signup</h2>
      <input placeholder="Email" onChange={e => setEmail(e.target.value)} /><br /><br />
      <input placeholder="Mobile" onChange={e => setMobile(e.target.value)} /><br /><br />
      <input type="password" placeholder="Password" onChange={e => setPassword(e.target.value)} /><br /><br />
      <button onClick={submit}>Signup</button>
    </div>
  );
}
