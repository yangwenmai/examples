#!/usr/bin/env python
# -*- coding: utf-8 -*-
import base64
import hashlib
import hmac

# In Flask and Django, raw_data is request.data.

# Example:
# if MTSigner(request.data, secrect_key) !=  request.headers.get('Authorization')
#	return {"Authorization is invalid"}, status.HTTP_401_BAD_REQUEST


class MTSigner(object):

    def __init__(self, raw_data, key):
        self.raw_data = raw_data
        self.key = key

    def sign(self):
        hash_str = hmac.new(self.key, self.raw_data, hashlib.sha1).hexdigest()
        sign_str = base64.b64encode(hash_str)
        return "meiqia_sign:" + sign_str
