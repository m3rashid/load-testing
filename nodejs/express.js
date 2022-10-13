const express = require("express");
const app = express();

app.use(express.urlencoded({ extended: true }));
app.use(express.json());

let hitCount = 0;

app.get("/", (req, res) => {
  hitCount += 1;

  const healthcheck = {
    uptime: process.uptime(),
    message: "OK",
    timestamp: Date.now(),
  };
  console.log("HIT:", hitCount, healthcheck);
  console.log();
  try {
    return res.send(healthcheck);
  } catch (error) {
    healthcheck.message = error;
    return res.status(503).send();
  }
});

app.listen(3000, () => console.log("Running on port 3000"));
