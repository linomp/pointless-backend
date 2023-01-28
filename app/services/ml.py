import os

import torch
from diffusers import StableDiffusionPipeline
from PIL.Image import Image
from dotenv import load_dotenv

load_dotenv()

token = os.getenv("HUGGINGFACE_TOKEN")

if token is None:
    raise ValueError("HUGGINGFACE_TOKEN is not set")

# get your token at https://huggingface.co/settings/tokens
pipe = StableDiffusionPipeline.from_pretrained(
    "CompVis/stable-diffusion-v1-4",
    revision="fp16",
    torch_dtype=torch.float16,
    use_auth_token=token,
)

pipe.to("cuda")


def obtain_image(
    prompt: str,
    *,
    seed: int | None = None,
    num_inference_steps: int = 5,
    guidance_scale: float = 7.5,
) -> Image:
    generator = None if seed is None else torch.Generator("cuda").manual_seed(seed)
    print(f"Using device: {pipe.device}")
    image: Image = pipe(
        prompt,
        guidance_scale=guidance_scale,
        num_inference_steps=num_inference_steps,
        generator=generator,
    ).images[0]
    return image

