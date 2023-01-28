from fastapi import APIRouter, Request
from fastapi.responses import HTMLResponse

from app.models.ServerMetrics import ServerMetrics
from app.utils.html import generate_server_status_html
from app.services.system import get_status

router = APIRouter()

tags = ["healthcheck"]


@router.get("/health/page", tags=tags)
async def get_healthcheck_page(request: Request) -> HTMLResponse:
    metrics = get_status(request)
    return HTMLResponse(content=generate_server_status_html(metrics), status_code=200)


@router.get("/health", tags=tags)
async def get_health(request: Request) -> ServerMetrics:
    return get_status(request)
