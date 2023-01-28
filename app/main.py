from fastapi import FastAPI

from app.routers import healthcheck, experiments
from fastapi.responses import RedirectResponse

app = FastAPI(title="pointless-backend", description="Backend for my Pointless personal site")

app.include_router(healthcheck.router)
app.include_router(experiments.router)


@app.get("/", tags=["healthcheck"])
async def redirect_to_healthcheck_page():
    return RedirectResponse(url="/health/page")
