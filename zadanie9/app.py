import os
import subprocess

import uvicorn
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel


class UserPrompt(BaseModel):
    prompt: str


class ChatResponse(BaseModel):
    response: str


def query_llama2(prompt: str) -> str:
    result = subprocess.run(
        ['ollama', 'run', 'llama2'],
        input=prompt.encode('utf-8'),
        capture_output=True,
        check=True
    )
    return result.stdout.strip()


app = FastAPI()


@app.post('/chat', response_model=ChatResponse)
async def post_prompt_to_llama(request: UserPrompt):
    user_prompt = request.prompt

    if not user_prompt:
        raise HTTPException(status_code=400, detail="No message provided")
    
    try:
        llama2_response = query_llama2(user_prompt)
        return ChatResponse(response=llama2_response)

    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
    

if __name__ == '__main__':
    uvicorn.run(app, host='0.0.0.0', port=8000)
