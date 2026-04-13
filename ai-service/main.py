
from fastapi import FastAPI
import random

app = FastAPI()

@app.get("/")
def root():
    return {"status": "AI running"}

@app.post("/predict")
def predict():
    return {
        "fraud": random.choice([0,1]),
        "score": round(random.random(),2)
    }
