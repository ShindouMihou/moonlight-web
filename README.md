<div align="center">
   once the darkness floods the skies, the moon will shine once again 
</div>

##
### 〰️ Moonlight

Moonlight is a simply-text editor platform intended for authors to write their chapters cross-platform. It is built to make editing chapters cross-platform more 
free and cost-effective by enabling self-hosting. This was created to be a sort of alternative to Pure Writer which lacked more functionality for desktop 
users.

#### Structure
- `midnight`: the server module of moonlight which is written in Golang and is backed by MongoDb.
- `sunshine`: the web client module of moonlight which is written in Svelte and Tailwind.

#### Authentication & Security
Moonlight is not intended to be a multi-user application, therefore, there is only one user that can be created and used and that is what we call the Root 
user which can be configured in Midnight. 

Authentication is handled simply through a JWT token that expires in 30-days, again, which may sound unsafe, but this project is clearly intended to be 
for a single person use and should be used in local network if not backed by a service such as Cloudflare Zero Trust.
