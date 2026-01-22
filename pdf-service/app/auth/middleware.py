from fastapi import Request
from starlette.responses import Response
from fastapi.security import HTTPBearer, HTTPAuthorizationCredentials
from .jwt_verifier import verify_jwt
from .exceptions import InvalidTokenException

security = HTTPBearer(auto_error=False)


async def auth_middleware(request: Request, call_next):
    credentials: HTTPAuthorizationCredentials | None = await security(request)

    if credentials:
        token = credentials.credentials
        try:
            request.state.user = verify_jwt(token)
        except InvalidTokenException:
            return Response("Invalid token", status_code=401)
    else:
        request.state.user = None

    response = await call_next(request)
    return response
