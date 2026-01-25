from jose import jwt, JWTError
from pathlib import Path
from .exceptions import InvalidTokenException

PUBLIC_KEY_PATH = Path(__file__).parent / "keys" / "public.pem"
with open(PUBLIC_KEY_PATH, "r") as f:
    PUBLIC_KEY = f.read()


ALGORITHM = "RS256"
ISSUER = "study-stack"


def verify_jwt(token: str) -> dict:
    """
    Verifies the JWT using RS256 and returns the claims when the token is
    valid.
    Raises InvalidTokenException if verification fails.
    """
    try:
        payload = jwt.decode(
            token,
            PUBLIC_KEY,
            algorithms=[ALGORITHM],
            options={"verify_exp": True},
            issuer=ISSUER
        )
        return payload
    except JWTError as e:
        raise InvalidTokenException() from e
