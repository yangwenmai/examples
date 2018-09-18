require 'base64'
require 'openssl'

# In rack apps that include Rails and Sinatra,
# raw_data is request.raw_body which only can be read once.

# Example:
# error!("Authorization is invalid", 401) unless MeiqiaToolsSigner.get(request.raw_body, secrect_key) == headers["Authorization"]

class MeiqiaToolsSigner
  def self.get(raw_body, key)
    hash_str = OpenSSL::HMAC.hexdigest(OpenSSL::Digest.new('sha1'), key, raw_body)
    sign_str = Base64.urlsafe_encode64(hash_str)
    "meiqia_sign:#{sign_str}"
  end
end
