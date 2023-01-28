import io

from fastapi import APIRouter
from fastapi.responses import StreamingResponse

from app.services.ml import obtain_image


router = APIRouter()

tags = ["experiments"]


@router.get("/image", tags=tags)
def generate_image_from_prompt(
        prompt: str = "photograph of an european short-hair cat",
        *,
        seed: int | None = None,
        num_inference_steps: int = 5,
        guidance_scale: float = 7.5
):
    image = obtain_image(
        prompt,
        num_inference_steps=num_inference_steps,
        seed=seed,
        guidance_scale=guidance_scale,
    )
    memory_stream = io.BytesIO()
    image.save(memory_stream, format="PNG")
    memory_stream.seek(0)
    return StreamingResponse(memory_stream, media_type="image/png")
