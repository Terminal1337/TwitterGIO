import random
import secrets
import binascii
import base64
import hashlib
from oauthlib import oauth1
import hmac
import collections
import string
from urllib.parse import urlencode
import time
from urllib.parse import quote
import sys


def randomString(length):
    """Returns random string"""
    return ''.join(secrets.choice(string.ascii_lowercase + string.ascii_lowercase) for i in range(length))

def escape(s):
    """Percent Encode the passed in string"""
    return quote(s, safe='~')

def generate_nonce():
    """Generate pseudorandom number."""

    randomByties = b"\x00" + secrets.token_bytes(32) + b"\x00"
    Baseencoded = base64.b64encode(randomByties)

    return Baseencoded.decode()


def stringify_parameters(parameters):
    """Orders parameters, and generates string representation of parameters"""
    output = ''
    ordered_parameters = {}
    ordered_parameters = collections.OrderedDict(sorted(parameters.items()))

    counter = 1
    for k, v in ordered_parameters.items():
        output += escape(str(k)) + '=' + escape(str(v))
        if counter < len(ordered_parameters):
            output += '&'
            counter += 1

    return output
def merge_two_dicts(x, y):
    z = x.copy()   # start with keys and values of x
    z.update(y)    # modifies z with keys and values of y
    return z
def SignatureString(data,url,method,payload):

    z = merge_two_dicts(data, payload)

    signature_base_string = (
            method + '&' +
            escape(url) + '&' +
            escape(stringify_parameters(z))
    )


    return signature_base_string

def SigningKey(secret):

    signingkeys = escape("GgDYlkSvaPxGxC4X8liwpUoqKwwr3lCADbz8A7ADU") + '&'
    signingkeys += escape(secret)

    return signingkeys


def calculate_signature(signing_key, signature_base_string):
    """Calculate the signature using SHA1"""
    hashed = hmac.new(signing_key.encode('utf-8'),signature_base_string.encode('utf-8'), hashlib.sha1)

    sig = binascii.b2a_base64(hashed.digest())[:-1]

    return escape(sig)

def getAuth(method,url,secret,token,payload):

    #initialize CLIENT FOR SIGNING
    client = oauth1.Client('3nVuSoBZnx6U4vzUxf5w',
                           client_secret='Bcs59EFbbsdF6Sl9Ng71smgStWEGwXXKSjYvPVt7qys',
                           resource_owner_key=token,
                           resource_owner_secret=secret,
                           signature_method=oauth1.SIGNATURE_HMAC_SHA1)

    if payload == None:
        payload = {
            'ext': 'highlightedLabel,mediaColor',
            'include_entities': '1',
            'include_profile_interstitial_type': 'true',
            'include_profile_location': 'true',
            'include_user_entities': 'true',
            'include_user_hashtag_entities': 'true',
            'include_user_mention_entities': 'true',
            'include_user_symbol_entities': 'true',
        }
        enc = urlencode(payload)
        headers = {'Content-Type': 'application/x-www-form-urlencoded'}
        uri, headers, body = client.sign(url,http_method=method,headers=headers,body=enc)
        return headers['Authorization']
    if payload == "NO_VALUE":
        uri, headers, body = client.sign(url, http_method=method)
        return headers['Authorization']
    else:
        enc = urlencode(payload)
        headers = {'Content-Type': 'application/x-www-form-urlencoded'}
        uri, headers, body = client.sign(url, http_method=method, headers=headers, body=enc)
        return headers['Authorization']



def getAuth2(method,url,secret,token,payload):
    auths = {
        'oauth_nonce': randomString(64),
        'oauth_timestamp': str(int(time.time())),
        'oauth_consumer_key': 'IQKbtAYlXLripLGPWd0HUA',
        'oauth_token': token,
        'oauth_version': '1.0',
        'oauth_signature_method': 'HMAC-SHA1',
    }

    if payload == None:
        payload = {
            'ext': 'highlightedLabel,mediaColor',
            'include_entities': '1',
            'include_profile_interstitial_type': 'true',
            'include_profile_location': 'true',
            'include_user_entities': 'true',
            'include_user_hashtag_entities': 'true',
            'include_user_mention_entities': 'true',
            'include_user_symbol_entities': 'true',
        }
    else:
        payload = payload
    signature_string = SignatureString(auths,url,method,payload)
    signing_key = SigningKey(secret)

    auths['oauth_signature'] = calculate_signature(signing_key,signature_string)

    ordered_parameters = {}
    ordered_parameters = collections.OrderedDict(sorted(auths.items()))
    auth_header = (
        '%s="%s"' % (k, v) for k, v in ordered_parameters.items())

    val = "OAuth " + ', '.join(auth_header)

    return val
token , secret = sys.argv[1],sys.argv[2]
print(getAuth("GET", 'https://twitter.com/account/authenticate_web_view', secret, token, "NO_VALUE"))