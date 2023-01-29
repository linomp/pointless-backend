import os

from fastapi import FastAPI

from fastapi.responses import RedirectResponse
from fastapi.middleware.cors import CORSMiddleware

from app.routers import healthcheck, experiments
app = FastAPI(title="pointless-backend", description="Backend for my Pointless personal site")

app.include_router(healthcheck.router)
app.include_router(experiments.router)

origins = [
    "https://pointless.xmp.systems",
    "http://localhost:3000",
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.get("/", tags=["healthcheck"])
async def redirect_to_healthcheck_page():
    return RedirectResponse(url="/health/page")
