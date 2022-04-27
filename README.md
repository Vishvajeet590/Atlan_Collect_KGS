# Atlan_Collect_KGS
Steps to run KGS

1. Clone this repo

        git clone https://github.com/Vishvajeet590/Atlan_Collect_KGS.git

2. Build the image

        cd ..
         docker build -t kgs -f Dockerfile.kgs ./

    This will take a few minutes.
 
3. Run the image's default command, which should start everything up. The `-p` option forwards the container's port 8080 to port 8000 on the host aong with the postgres DB url as env variable. (Note that the host will actually be a guest if you are using boot2docker, so you may need to re-forward the port in VirtualBox.)

        docker run --rm -p 8080:8080 -e KEY_DATABASE_URL='postgres://vishwajeet:docker@localhost:5431/KeyStore-1?&pool_max_conns=10' --network="host" kgs

    
    
