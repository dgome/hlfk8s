# Hyperledger Fabric & Kubernetes

_Under Development_

## Create Namespace

```sh
kubectl delete -f namespace.yaml

kubectl apply -f namespace.yaml
```

## Create Volume

```sh
kubectl apply -f pv1.yaml
```

## Create Persistent Volume Claim

```sh
kubectl apply -f pvc1.yaml
```

## Run Fabric Tools

```sh
kubectl apply -f fabric-tools.yaml
```

## Run CAs

```sh
kubectl apply -f orderers-ca.yaml
kubectl apply -f org1-ca.yaml
kubectl apply -f org2-ca.yaml
```

## Orderers Org Registers & Enrolls

### Enroll Orderers Bootstrap Identity

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/msp
kubectl exec -n hyperledger deploy/orderers-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/fabric-ca-client-config.yaml

kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client enroll --url https://admin:adminpw@localhost:7054

kubectl cp config/config.yaml hyperledger/fabric-tools:/vol1/organizations/ordererOrganizations/orderers/msp/
```

### Register Admin1

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client register --id.name admin1 --id.secret admin1pw --id.type admin --url https://admin:adminpw@localhost:7054
```

### Enroll Admin1

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/users/admin1/msp

kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client enroll --url https://admin1:admin1pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/users/admin1/msp/

kubectl cp config/config.yaml hyperledger/fabric-tools:/vol1/organizations/ordererOrganizations/orderers/users/admin1/msp/
```

### Register Orderer0

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client register --id.name orderer0 --id.secret orderer0pw --id.type orderer --url https://admin:adminpw@localhost:7054
```

### Enroll Orderer0

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/orderers/orderer0/msp

kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client enroll --url https://orderer0:orderer0pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/orderers/orderer0/msp/

kubectl cp config/config.yaml hyperledger/fabric-tools:/vol1/organizations/ordererOrganizations/orderers/orderers/orderer0/msp/
```

### TLS Enroll Orderer0

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/orderers/orderer0/tls

kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client enroll --url https://orderer0:orderer0pw@localhost:7054 --enrollment.profile tls --csr.hosts localhost --csr.hosts orderer0 --mspdir /etc/hyperledger/fabric-ca-client/orderers/orderer0/tls/

kubectl exec -n hyperledger deploy/orderers-ca -it -- bash -c 'cp /etc/hyperledger/fabric-ca-client/orderers/orderer0/tls/keystore/* /etc/hyperledger/fabric-ca-client/orderers/orderer0/tls/server.key'

kubectl exec -n hyperledger deploy/orderers-ca -it -- mkdir -p /etc/hyperledger/fabric-ca-client/msp/tlscacerts/

kubectl exec -n hyperledger deploy/orderers-ca -it -- cp /etc/hyperledger/fabric-ca-client/orderers/orderer0/tls/tlscacerts/tls-localhost-7054.pem /etc/hyperledger/fabric-ca-client/msp/tlscacerts/
```

### Register Orderer1

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client register --id.name orderer1 --id.secret orderer1pw --id.type orderer --url https://admin:adminpw@localhost:7054
```

### Enroll Orderer1

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/orderers/orderer1/msp

kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client enroll --url https://orderer1:orderer1pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/orderers/orderer1/msp/

kubectl cp config/config.yaml hyperledger/fabric-tools:/vol1/organizations/ordererOrganizations/orderers/orderers/orderer1/msp/
```

### TLS Enroll Orderer1

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/orderers/orderer1/tls

kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client enroll --url https://orderer1:orderer1pw@localhost:7054 --enrollment.profile tls --csr.hosts localhost --csr.hosts orderer1 --mspdir /etc/hyperledger/fabric-ca-client/orderers/orderer1/tls/

kubectl exec -n hyperledger deploy/orderers-ca -it -- bash -c 'cp /etc/hyperledger/fabric-ca-client/orderers/orderer1/tls/keystore/* /etc/hyperledger/fabric-ca-client/orderers/orderer1/tls/server.key'
```

### Register Orderer2

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client register --id.name orderer2 --id.secret orderer2pw --id.type orderer --url https://admin:adminpw@localhost:7054
```

### Enroll Orderer2

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/orderers/orderer2/msp

kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client enroll --url https://orderer2:orderer2pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/orderers/orderer2/msp/

kubectl cp config/config.yaml hyperledger/fabric-tools:/vol1/organizations/ordererOrganizations/orderers/orderers/orderer2/msp/
```

### TLS Enroll Orderer2

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/orderers/orderer2/tls

kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client enroll --url https://orderer2:orderer2pw@localhost:7054 --enrollment.profile tls --csr.hosts localhost --csr.hosts orderer2 --mspdir /etc/hyperledger/fabric-ca-client/orderers/orderer2/tls/

kubectl exec -n hyperledger deploy/orderers-ca -it -- bash -c 'cp /etc/hyperledger/fabric-ca-client/orderers/orderer2/tls/keystore/* /etc/hyperledger/fabric-ca-client/orderers/orderer2/tls/server.key'
```

## Org1 Registers & Enrolls

### Enroll Org1 Bootstrap Identity

```sh
kubectl exec -n hyperledger deploy/org1-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/msp
kubectl exec -n hyperledger deploy/org1-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/fabric-ca-client-config.yaml

kubectl exec -n hyperledger deploy/org1-ca -it -- fabric-ca-client enroll --url https://admin:adminpw@localhost:7054

kubectl cp config/config.yaml hyperledger/fabric-tools:/vol1/organizations/peerOrganizations/org1/msp/
```

### Register Admin1

```sh
kubectl exec -n hyperledger deploy/org1-ca -it -- fabric-ca-client register --id.name admin1 --id.secret admin1pw --id.type admin --url https://admin:adminpw@localhost:7054
```

### Enroll Admin1

```sh
kubectl exec -n hyperledger deploy/org1-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/users/admin1/msp

kubectl exec -n hyperledger deploy/org1-ca -it -- fabric-ca-client enroll --url https://admin1:admin1pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/users/admin1/msp/

kubectl cp config/config.yaml hyperledger/fabric-tools:/vol1/organizations/peerOrganizations/org1/users/admin1/msp/
```

### Register Peer0

```sh
kubectl exec -n hyperledger deploy/org1-ca -it -- fabric-ca-client register --id.name peer0 --id.secret peer0pw --id.type peer --url https://admin:adminpw@localhost:7054
```

### Enroll Peer0

```sh
kubectl exec -n hyperledger deploy/org1-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/peers/peer0/msp

kubectl exec -n hyperledger deploy/org1-ca -it -- fabric-ca-client enroll --url https://peer0:peer0pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/peers/peer0/msp/

kubectl cp config/config.yaml hyperledger/fabric-tools:/vol1/organizations/peerOrganizations/org1/peers/peer0/msp/
```

### TLS Enroll Peer0

```sh
kubectl exec -n hyperledger deploy/org1-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/peers/peer0/tls

kubectl exec -n hyperledger deploy/org1-ca -it -- fabric-ca-client enroll --url https://peer0:peer0pw@localhost:7054 --enrollment.profile tls --csr.hosts localhost --csr.hosts org1-peer0 --mspdir /etc/hyperledger/fabric-ca-client/peers/peer0/tls/

kubectl exec -n hyperledger deploy/org1-ca -it -- bash -c 'cp /etc/hyperledger/fabric-ca-client/peers/peer0/tls/keystore/* /etc/hyperledger/fabric-ca-client/peers/peer0/tls/server.key'
```

### Register Peer1

```sh
kubectl exec -n hyperledger deploy/org1-ca -it -- fabric-ca-client register --id.name peer1 --id.secret peer1pw --id.type peer --url https://admin:adminpw@localhost:7054
```

### Enroll Peer1

```sh
kubectl exec -n hyperledger deploy/org1-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/peers/peer1/msp

kubectl exec -n hyperledger deploy/org1-ca -it -- fabric-ca-client enroll --url https://peer1:peer1pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/peers/peer1/msp/

kubectl cp config/config.yaml hyperledger/fabric-tools:/vol1/organizations/peerOrganizations/org1/peers/peer1/msp/
```

### TLS Enroll Peer1

```sh
kubectl exec -n hyperledger deploy/org1-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/peers/peer1/tls

kubectl exec -n hyperledger deploy/org1-ca -it -- fabric-ca-client enroll --url https://peer1:peer1pw@localhost:7054 --enrollment.profile tls --csr.hosts localhost --csr.hosts org1-peer1 --mspdir /etc/hyperledger/fabric-ca-client/peers/peer1/tls/

kubectl exec -n hyperledger deploy/org1-ca -it -- bash -c 'cp /etc/hyperledger/fabric-ca-client/peers/peer1/tls/keystore/* /etc/hyperledger/fabric-ca-client/peers/peer1/tls/server.key'
```

## Org2 Registers & Enrolls

### Enroll Org2 Bootstrap Identity

```sh
kubectl exec -n hyperledger deploy/org2-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/msp
kubectl exec -n hyperledger deploy/org2-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/fabric-ca-client-config.yaml

kubectl exec -n hyperledger deploy/org2-ca -it -- fabric-ca-client enroll --url https://admin:adminpw@localhost:7054

kubectl cp config/config.yaml hyperledger/fabric-tools:/vol1/organizations/peerOrganizations/org2/msp/
```

### Register Admin1

```sh
kubectl exec -n hyperledger deploy/org2-ca -it -- fabric-ca-client register --id.name admin1 --id.secret admin1pw --id.type admin --url https://admin:adminpw@localhost:7054
```

### Enroll Admin1

```sh
kubectl exec -n hyperledger deploy/org2-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/users/admin1/msp

kubectl exec -n hyperledger deploy/org2-ca -it -- fabric-ca-client enroll --url https://admin1:admin1pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/users/admin1/msp/

kubectl cp config/config.yaml hyperledger/fabric-tools:/vol1/organizations/peerOrganizations/org2/users/admin1/msp/
```

### Register Peer0

```sh
kubectl exec -n hyperledger deploy/org2-ca -it -- fabric-ca-client register --id.name peer0 --id.secret peer0pw --id.type peer --url https://admin:adminpw@localhost:7054
```

### Enroll Peer0

```sh
kubectl exec -n hyperledger deploy/org2-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/peers/peer0/msp

kubectl exec -n hyperledger deploy/org2-ca -it -- fabric-ca-client enroll --url https://peer0:peer0pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/peers/peer0/msp/

kubectl cp config/config.yaml hyperledger/fabric-tools:/vol1/organizations/peerOrganizations/org2/peers/peer0/msp/
```

### TLS Enroll Peer0

```sh
kubectl exec -n hyperledger deploy/org2-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/peers/peer0/tls

kubectl exec -n hyperledger deploy/org2-ca -it -- fabric-ca-client enroll --url https://peer0:peer0pw@localhost:7054 --enrollment.profile tls --csr.hosts localhost --csr.hosts org2-peer0 --mspdir /etc/hyperledger/fabric-ca-client/peers/peer0/tls/

kubectl exec -n hyperledger deploy/org2-ca -it -- bash -c 'cp /etc/hyperledger/fabric-ca-client/peers/peer0/tls/keystore/* /etc/hyperledger/fabric-ca-client/peers/peer0/tls/server.key'
```

### Register Peer1

```sh
kubectl exec -n hyperledger deploy/org2-ca -it -- fabric-ca-client register --id.name peer1 --id.secret peer1pw --id.type peer --url https://admin:adminpw@localhost:7054
```

### Enroll Peer1

```sh
kubectl exec -n hyperledger deploy/org2-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/peers/peer1/msp

kubectl exec -n hyperledger deploy/org2-ca -it -- fabric-ca-client enroll --url https://peer1:peer1pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/peers/peer1/msp/

kubectl cp config/config.yaml hyperledger/fabric-tools:/vol1/organizations/peerOrganizations/org2/peers/peer1/msp/
```

### TLS Enroll Peer1

```sh
kubectl exec -n hyperledger deploy/org2-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/peers/peer1/tls

kubectl exec -n hyperledger deploy/org2-ca -it -- fabric-ca-client enroll --url https://peer1:peer1pw@localhost:7054 --enrollment.profile tls --csr.hosts localhost --csr.hosts org2-peer1 --mspdir /etc/hyperledger/fabric-ca-client/peers/peer1/tls/

kubectl exec -n hyperledger deploy/org2-ca -it -- bash -c 'cp /etc/hyperledger/fabric-ca-client/peers/peer1/tls/keystore/* /etc/hyperledger/fabric-ca-client/peers/peer1/tls/server.key'
```

## Copy `configtx.yaml`

```sh
kubectl exec -n hyperledger fabric-tools -it -- mkdir -p /vol1/config/

kubectl cp config/configtx.yaml hyperledger/fabric-tools:/vol1/config/
```

## Create Genesis Block

```sh
kubectl exec -n hyperledger fabric-tools -it -- configtxgen -configPath /vol1/config/ -profile OrdererGenesis -channelID syschannel -outputBlock /vol1/genesis.block
```

## Run Orderers

```sh
kubectl apply -f orderer0.yaml
kubectl apply -f orderer1.yaml
kubectl apply -f orderer2.yaml
```

## Run Peers

### Org1

```sh
kubectl apply -f org1-peer0.yaml
kubectl apply -f org1-peer1.yaml
```

### Org2

```sh
kubectl apply -f org2-peer0.yaml
kubectl apply -f org2-peer1.yaml
```

## Create Channel Transaction

```sh
kubectl exec -n hyperledger fabric-tools -it -- configtxgen -configPath /vol1/config/ -profile MyChannel -outputCreateChannelTx /vol1/channel-artifacts/mychannel.tx -channelID mychannel
```

## Create Org Anchors Transaction

### Org1

```sh
kubectl exec -n hyperledger fabric-tools -it -- configtxgen -configPath /vol1/config/ -profile MyChannel -outputAnchorPeersUpdate /vol1/channel-artifacts/Org1-Anchors.tx -channelID mychannel -asOrg Org1
```

### Org2

```sh
kubectl exec -n hyperledger fabric-tools -it -- configtxgen -configPath /vol1/config/ -profile MyChannel -outputAnchorPeersUpdate /vol1/channel-artifacts/Org2-Anchors.tx -channelID mychannel -asOrg Org2
```

## Create Channel by Org1 Peer0

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c 'export CORE_PEER_LOCALMSPID=Org1MSP; export CORE_PEER_MSPCONFIGPATH=/vol1/organizations/peerOrganizations/org1/users/admin1/msp; peer channel create --channelID mychannel --file /vol1/channel-artifacts/mychannel.tx --outputBlock /vol1/channel-artifacts/mychannel.block --tls --cafile /vol1/organizations/ordererOrganizations/orderers/msp/tlscacerts/tls-localhost-7054.pem --orderer orderer0:7050 --ordererTLSHostnameOverride orderer0'
```

## Join Channel

### Org1

#### Peer 0

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c '\
export CORE_PEER_LOCALMSPID="Org1MSP"; \
export CORE_PEER_MSPCONFIGPATH=/vol1/organizations/peerOrganizations/org1/users/admin1/msp; \
export CORE_PEER_ADDRESS=org1-peer0:7051; \
export CORE_PEER_TLS_ENABLED=true; \
export CORE_PEER_TLS_ROOTCERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/tlscacerts/tls-localhost-7054.pem; \
peer channel join --blockpath /vol1/channel-artifacts/mychannel.block'
```

#### Peer 1

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c '\
export CORE_PEER_LOCALMSPID="Org1MSP"; \
export CORE_PEER_MSPCONFIGPATH=/vol1/organizations/peerOrganizations/org1/users/admin1/msp; \
export CORE_PEER_ADDRESS=org1-peer1:7051; \
export CORE_PEER_TLS_ENABLED=true; \
export CORE_PEER_TLS_ROOTCERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer1/tls/tlscacerts/tls-localhost-7054.pem; \
peer channel join --blockpath /vol1/channel-artifacts/mychannel.block'
```

### Org2

#### Peer 0

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c '\
export CORE_PEER_LOCALMSPID="Org2MSP"; \
export CORE_PEER_MSPCONFIGPATH=/vol1/organizations/peerOrganizations/org2/users/admin1/msp; \
export CORE_PEER_ADDRESS=org2-peer0:7051; \
export CORE_PEER_TLS_ENABLED=true; \
export CORE_PEER_TLS_ROOTCERT_FILE=/vol1/organizations/peerOrganizations/org2/peers/peer0/tls/tlscacerts/tls-localhost-7054.pem; \
peer channel join --blockpath /vol1/channel-artifacts/mychannel.block'
```

#### Peer 1

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c '\
export CORE_PEER_LOCALMSPID="Org2MSP"; \
export CORE_PEER_MSPCONFIGPATH=/vol1/organizations/peerOrganizations/org2/users/admin1/msp; \
export CORE_PEER_ADDRESS=org2-peer1:7051; \
export CORE_PEER_TLS_ENABLED=true; \
export CORE_PEER_TLS_ROOTCERT_FILE=/vol1/organizations/peerOrganizations/org2/peers/peer1/tls/tlscacerts/tls-localhost-7054.pem; \
peer channel join --blockpath /vol1/channel-artifacts/mychannel.block'
```

## Update Anchor Peers

### Org1

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c '\
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=/vol1/organizations/peerOrganizations/org1/users/admin1/msp; \
export CORE_PEER_ADDRESS=org1-peer0:7051; \
export CORE_PEER_TLS_ROOTCERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/tlscacerts/tls-localhost-7054.pem; \
peer channel update --orderer orderer0:7050 --ordererTLSHostnameOverride orderer0 --channelID mychannel --file /vol1/channel-artifacts/Org1-Anchors.tx --tls true --cafile /vol1/organizations/ordererOrganizations/orderers/msp/tlscacerts/tls-localhost-7054.pem'
```

### Org2

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c '\
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_MSPCONFIGPATH=/vol1/organizations/peerOrganizations/org2/users/admin1/msp; \
export CORE_PEER_ADDRESS=org2-peer0:7051; \
export CORE_PEER_TLS_ROOTCERT_FILE=/vol1/organizations/peerOrganizations/org2/peers/peer0/tls/tlscacerts/tls-localhost-7054.pem; \
peer channel update --orderer orderer0:7050 --ordererTLSHostnameOverride orderer0 --channelID mychannel --file /vol1/channel-artifacts/Org2-Anchors.tx --tls true --cafile /vol1/organizations/ordererOrganizations/orderers/msp/tlscacerts/tls-localhost-7054.pem'
```

## Install Chain Code

### Copy

```sh
kubectl -n hyperledger exec fabric-tools -it -- rm -rf /vol1/chaincode0
kubectl cp chaincode0 hyperledger/fabric-tools:/vol1/chaincode0
```

### Vendor

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c 'pushd /vol1/chaincode0; GO111MODULE=on go mod vendor'
```

### Package

```sh
kubectl -n hyperledger exec fabric-tools -it -- rm -f /vol1/chaincode0.tar.gz
kubectl exec -n hyperledger fabric-tools -it -- peer lifecycle chaincode package /vol1/chaincode0.tar.gz --path /vol1/chaincode0 --lang golang --label cc0_v1
```

### Install

#### Org1

##### Peer0

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c '\
export CORE_PEER_LOCALMSPID="Org1MSP"; \
export CORE_PEER_MSPCONFIGPATH=/vol1/organizations/peerOrganizations/org1/users/admin1/msp; \
export CORE_PEER_ADDRESS=org1-peer0:7051; \
export CORE_PEER_TLS_ENABLED=true; \
export CORE_PEER_TLS_ROOTCERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/tlscacerts/tls-localhost-7054.pem; \
export CORE_PEER_TLS_CERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/signcerts/cert.pem; \
export CORE_PEER_TLS_KEY_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/server.key; \
peer lifecycle chaincode install /vol1/chaincode0.tar.gz'
```

##### Peer1

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c '\
export CORE_PEER_LOCALMSPID="Org1MSP"; \
export CORE_PEER_MSPCONFIGPATH=/vol1/organizations/peerOrganizations/org1/users/admin1/msp; \
export CORE_PEER_ADDRESS=org1-peer1:7051; \
export CORE_PEER_TLS_ENABLED=true; \
export CORE_PEER_TLS_ROOTCERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer1/tls/tlscacerts/tls-localhost-7054.pem; \
export CORE_PEER_TLS_CERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer1/tls/signcerts/cert.pem; \
export CORE_PEER_TLS_KEY_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer1/tls/server.key; \
peer lifecycle chaincode install /vol1/chaincode0.tar.gz'
```

#### Org2

##### Peer0

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c '\
export CORE_PEER_LOCALMSPID="Org2MSP"; \
export CORE_PEER_MSPCONFIGPATH=/vol1/organizations/peerOrganizations/org2/users/admin1/msp; \
export CORE_PEER_ADDRESS=org2-peer0:7051; \
export CORE_PEER_TLS_ENABLED=true; \
export CORE_PEER_TLS_ROOTCERT_FILE=/vol1/organizations/peerOrganizations/org2/peers/peer0/tls/tlscacerts/tls-localhost-7054.pem; \
export CORE_PEER_TLS_CERT_FILE=/vol1/organizations/peerOrganizations/org2/peers/peer0/tls/signcerts/cert.pem; \
export CORE_PEER_TLS_KEY_FILE=/vol1/organizations/peerOrganizations/org2/peers/peer0/tls/server.key; \
peer lifecycle chaincode install /vol1/chaincode0.tar.gz'
```

##### Peer1

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c '\
export CORE_PEER_LOCALMSPID="Org2MSP"; \
export CORE_PEER_MSPCONFIGPATH=/vol1/organizations/peerOrganizations/org2/users/admin1/msp; \
export CORE_PEER_ADDRESS=org2-peer1:7051; \
export CORE_PEER_TLS_ENABLED=true; \
export CORE_PEER_TLS_ROOTCERT_FILE=/vol1/organizations/peerOrganizations/org2/peers/peer1/tls/tlscacerts/tls-localhost-7054.pem; \
export CORE_PEER_TLS_CERT_FILE=/vol1/organizations/peerOrganizations/org2/peers/peer1/tls/signcerts/cert.pem; \
export CORE_PEER_TLS_KEY_FILE=/vol1/organizations/peerOrganizations/org2/peers/peer1/tls/server.key; \
peer lifecycle chaincode install /vol1/chaincode0.tar.gz'
```

#### Query Installed Chain Codes

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c '\
export CORE_PEER_LOCALMSPID="Org1MSP"; \
export CORE_PEER_MSPCONFIGPATH=/vol1/organizations/peerOrganizations/org1/users/admin1/msp; \
export CORE_PEER_ADDRESS=org1-peer0:7051; \
export CORE_PEER_TLS_ENABLED=true; \
export CORE_PEER_TLS_ROOTCERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/tlscacerts/tls-localhost-7054.pem; \
export CORE_PEER_TLS_CERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/signcerts/cert.pem; \
export CORE_PEER_TLS_KEY_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/server.key; \
peer lifecycle chaincode queryinstalled'
```

### Approve

#### Org1 by Peer0

_TODO: Auto Get Package ID_

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c '\
export CORE_PEER_LOCALMSPID="Org1MSP"; \
export CORE_PEER_MSPCONFIGPATH=/vol1/organizations/peerOrganizations/org1/users/admin1/msp; \
export CORE_PEER_ADDRESS=org1-peer0:7051; \
export CORE_PEER_TLS_ENABLED=true; \
export CORE_PEER_TLS_ROOTCERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/tlscacerts/tls-localhost-7054.pem; \
export CORE_PEER_TLS_CERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/signcerts/cert.pem; \
export CORE_PEER_TLS_KEY_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/server.key; \
peer lifecycle chaincode approveformyorg --orderer orderer0:7050 --ordererTLSHostnameOverride orderer0 --tls true --cafile /vol1/organizations/ordererOrganizations/orderers/msp/tlscacerts/tls-localhost-7054.pem --channelID mychannel --name cc0 --version 1 --init-required --package-id cc0_v1:90c598130fc773cc579287d52e62fbaee0b2ffe53874d19a244b6bbb98f32fad --sequence 1'
```

#### Org2 by Peer0

_TODO: Auto Get Package ID_

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c '\
export CORE_PEER_LOCALMSPID="Org2MSP"; \
export CORE_PEER_MSPCONFIGPATH=/vol1/organizations/peerOrganizations/org2/users/admin1/msp; \
export CORE_PEER_ADDRESS=org2-peer0:7051; \
export CORE_PEER_TLS_ENABLED=true; \
export CORE_PEER_TLS_ROOTCERT_FILE=/vol1/organizations/peerOrganizations/org2/peers/peer0/tls/tlscacerts/tls-localhost-7054.pem; \
export CORE_PEER_TLS_CERT_FILE=/vol1/organizations/peerOrganizations/org2/peers/peer0/tls/signcerts/cert.pem; \
export CORE_PEER_TLS_KEY_FILE=/vol1/organizations/peerOrganizations/org2/peers/peer0/tls/server.key; \
peer lifecycle chaincode approveformyorg --orderer orderer0:7050 --ordererTLSHostnameOverride orderer0 --tls true --cafile /vol1/organizations/ordererOrganizations/orderers/msp/tlscacerts/tls-localhost-7054.pem --channelID mychannel --name cc0 --version 1 --init-required --package-id cc0_v1:90c598130fc773cc579287d52e62fbaee0b2ffe53874d19a244b6bbb98f32fad --sequence 1'
```

### Commit Chaincode Definition

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c '\
export CORE_PEER_LOCALMSPID="Org1MSP"; \
export CORE_PEER_MSPCONFIGPATH=/vol1/organizations/peerOrganizations/org1/users/admin1/msp; \
export CORE_PEER_ADDRESS=org1-peer0:7051; \
export CORE_PEER_TLS_ENABLED=true; \
export CORE_PEER_TLS_ROOTCERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/tlscacerts/tls-localhost-7054.pem; \
export CORE_PEER_TLS_CERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/signcerts/cert.pem; \
export CORE_PEER_TLS_KEY_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/server.key; \
peer lifecycle chaincode commit --orderer orderer0:7050 --ordererTLSHostnameOverride orderer0 --tls true --cafile /vol1/organizations/ordererOrganizations/orderers/msp/tlscacerts/tls-localhost-7054.pem --channelID mychannel --name cc0 --version 1 --sequence 1 --init-required'
```

### Invoke Init

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c '\
export CORE_PEER_LOCALMSPID="Org1MSP"; \
export CORE_PEER_MSPCONFIGPATH=/vol1/organizations/peerOrganizations/org1/users/admin1/msp; \
export CORE_PEER_ADDRESS=org1-peer0:7051; \
export CORE_PEER_TLS_ENABLED=true; \
export CORE_PEER_TLS_ROOTCERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/tlscacerts/tls-localhost-7054.pem; \
export CORE_PEER_TLS_CERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/signcerts/cert.pem; \
export CORE_PEER_TLS_KEY_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/server.key; \
peer chaincode invoke --orderer orderer0:7050 --ordererTLSHostnameOverride orderer0 --tls true --cafile /vol1/organizations/ordererOrganizations/orderers/msp/tlscacerts/tls-localhost-7054.pem --channelID mychannel --name cc0 --peerAddresses org1-peer0:7051 --tlsRootCertFiles /vol1/organizations/peerOrganizations/org1/peers/peer0/tls/tlscacerts/tls-localhost-7054.pem --peerAddresses org2-peer0:7051 --tlsRootCertFiles /vol1/organizations/peerOrganizations/org2/peers/peer0/tls/tlscacerts/tls-localhost-7054.pem --isInit --ctor "{\"function\":\"InitLedger\",\"Args\":[]}"'
```

## Invoke Chain Code

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c '\
export CORE_PEER_LOCALMSPID="Org1MSP"; \
export CORE_PEER_MSPCONFIGPATH=/vol1/organizations/peerOrganizations/org1/users/admin1/msp; \
export CORE_PEER_ADDRESS=org1-peer0:7051; \
export CORE_PEER_TLS_ENABLED=true; \
export CORE_PEER_TLS_ROOTCERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/tlscacerts/tls-localhost-7054.pem; \
export CORE_PEER_TLS_CERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/signcerts/cert.pem; \
export CORE_PEER_TLS_KEY_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/server.key; \
peer chaincode invoke --orderer orderer0:7050 --ordererTLSHostnameOverride orderer0 --tls true --cafile /vol1/organizations/ordererOrganizations/orderers/msp/tlscacerts/tls-localhost-7054.pem --channelID mychannel --name cc0 --peerAddresses org1-peer0:7051 --tlsRootCertFiles /vol1/organizations/peerOrganizations/org1/peers/peer0/tls/tlscacerts/tls-localhost-7054.pem --peerAddresses org2-peer0:7051 --tlsRootCertFiles /vol1/organizations/peerOrganizations/org2/peers/peer0/tls/tlscacerts/tls-localhost-7054.pem --ctor "{\"Args\":[\"CreateItem\", \"1\", \"Item 1\"]}"'
```

## Query Chain Code

```sh
kubectl exec -n hyperledger fabric-tools -it -- bash -c '\
export CORE_PEER_LOCALMSPID="Org1MSP"; \
export CORE_PEER_MSPCONFIGPATH=/vol1/organizations/peerOrganizations/org1/users/admin1/msp; \
export CORE_PEER_ADDRESS=org1-peer0:7051; \
export CORE_PEER_TLS_ENABLED=true; \
export CORE_PEER_TLS_ROOTCERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/tlscacerts/tls-localhost-7054.pem; \
export CORE_PEER_TLS_CERT_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/signcerts/cert.pem; \
export CORE_PEER_TLS_KEY_FILE=/vol1/organizations/peerOrganizations/org1/peers/peer0/tls/server.key; \
peer chaincode query --channelID mychannel --name cc0 --ctor "{\"Args\":[\"GetItem\", \"1\"]}"'
```
