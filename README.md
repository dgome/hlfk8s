# Hyperledger Fabric & Kubernetes

_Under Development_

## Create Namespace

```sh
kubectl apply -f namespace.yaml
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

kubectl cp configs/config.yaml hyperledger/fabric-tools:/vol1/organizations/ordererOrganizations/orderers/msp/
```

### Register Admin1

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client register --id.name admin1 --id.secret admin1pw --id.type admin --url https://admin:adminpw@localhost:7054
```

### Enroll Admin1

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/users/admin1/msp

kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client enroll --url https://admin1:admin1pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/users/admin1/msp/

kubectl cp configs/config.yaml hyperledger/fabric-tools:/vol1/organizations/ordererOrganizations/orderers/users/admin1/msp/
```

### Register Orderer0

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client register --id.name orderer0 --id.secret orderer0pw --id.type orderer --url https://admin:adminpw@localhost:7054
```

### Enroll Orderer0

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/orderers/orderer0/msp

kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client enroll --url https://orderer0:orderer0pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/orderers/orderer0/msp/

kubectl cp configs/config.yaml hyperledger/fabric-tools:/vol1/organizations/ordererOrganizations/orderers/orderers/orderer0/msp/
```

### TLS Enroll Orderer0

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/orderers/orderer0/tls

kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client enroll --url https://orderer0:orderer0pw@localhost:7054 --enrollment.profile tls --csr.hosts localhost --csr.hosts orderer0 --mspdir /etc/hyperledger/fabric-ca-client/orderers/orderer0/tls/

kubectl exec -n hyperledger deploy/orderers-ca -it -- bash -c 'cp /etc/hyperledger/fabric-ca-client/orderers/orderer0/tls/keystore/* /etc/hyperledger/fabric-ca-client/orderers/orderer0/tls/server.key'
```

### Register Orderer1

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client register --id.name orderer1 --id.secret orderer1pw --id.type orderer --url https://admin:adminpw@localhost:7054
```

### Enroll Orderer1

```sh
kubectl exec -n hyperledger deploy/orderers-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/orderers/orderer1/msp

kubectl exec -n hyperledger deploy/orderers-ca -it -- fabric-ca-client enroll --url https://orderer1:orderer1pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/orderers/orderer1/msp/

kubectl cp configs/config.yaml hyperledger/fabric-tools:/vol1/organizations/ordererOrganizations/orderers/orderers/orderer1/msp/
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

kubectl cp configs/config.yaml hyperledger/fabric-tools:/vol1/organizations/ordererOrganizations/orderers/orderers/orderer2/msp/
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

kubectl cp configs/config.yaml hyperledger/fabric-tools:/vol1/organizations/peerOrganizations/org1/msp/
```

### Register Admin1

```sh
kubectl exec -n hyperledger deploy/org1-ca -it -- fabric-ca-client register --id.name admin1 --id.secret admin1pw --id.type admin --url https://admin:adminpw@localhost:7054
```

### Enroll Admin1

```sh
kubectl exec -n hyperledger deploy/org1-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/users/admin1/msp

kubectl exec -n hyperledger deploy/org1-ca -it -- fabric-ca-client enroll --url https://admin1:admin1pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/users/admin1/msp/

kubectl cp configs/config.yaml hyperledger/fabric-tools:/vol1/organizations/peerOrganizations/org1/users/admin1/msp/
```

### Register Peer0

```sh
kubectl exec -n hyperledger deploy/org1-ca -it -- fabric-ca-client register --id.name peer0 --id.secret peer0pw --id.type peer --url https://admin:adminpw@localhost:7054
```

### Enroll Peer0

```sh
kubectl exec -n hyperledger deploy/org1-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/peers/peer0/msp

kubectl exec -n hyperledger deploy/org1-ca -it -- fabric-ca-client enroll --url https://peer0:peer0pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/peers/peer0/msp/

kubectl cp configs/config.yaml hyperledger/fabric-tools:/vol1/organizations/peerOrganizations/org1/peers/peer0/msp/
```

### TLS Enroll Peer0

```sh
kubectl exec -n hyperledger deploy/org1-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/peers/peer0/tls

kubectl exec -n hyperledger deploy/org1-ca -it -- fabric-ca-client enroll --url https://peer0:peer0pw@localhost:7054 --enrollment.profile tls --csr.hosts localhost --csr.hosts peer0-org1 --mspdir /etc/hyperledger/fabric-ca-client/peers/peer0/tls/
```

### Register Peer1

```sh
kubectl exec -n hyperledger deploy/org1-ca -it -- fabric-ca-client register --id.name peer1 --id.secret peer1pw --id.type peer --url https://admin:adminpw@localhost:7054
```

### Enroll Peer1

```sh
kubectl exec -n hyperledger deploy/org1-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/peers/peer1/msp

kubectl exec -n hyperledger deploy/org1-ca -it -- fabric-ca-client enroll --url https://peer1:peer1pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/peers/peer1/msp/

kubectl cp configs/config.yaml hyperledger/fabric-tools:/vol1/organizations/peerOrganizations/org1/peers/peer1/msp/
```

### TLS Enroll Peer1

```sh
kubectl exec -n hyperledger deploy/org1-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/peers/peer1/tls

kubectl exec -n hyperledger deploy/org1-ca -it -- fabric-ca-client enroll --url https://peer1:peer1pw@localhost:7054 --enrollment.profile tls --csr.hosts localhost --csr.hosts peer1-org1 --mspdir /etc/hyperledger/fabric-ca-client/peers/peer1/tls/
```

## Org2 Registers & Enrolls

### Enroll Org2 Bootstrap Identity

```sh
kubectl exec -n hyperledger deploy/org2-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/msp
kubectl exec -n hyperledger deploy/org2-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/fabric-ca-client-config.yaml

kubectl exec -n hyperledger deploy/org2-ca -it -- fabric-ca-client enroll --url https://admin:adminpw@localhost:7054

kubectl cp configs/config.yaml hyperledger/fabric-tools:/vol1/organizations/peerOrganizations/org2/msp/
```

### Register Admin1

```sh
kubectl exec -n hyperledger deploy/org2-ca -it -- fabric-ca-client register --id.name admin1 --id.secret admin1pw --id.type admin --url https://admin:adminpw@localhost:7054
```

### Enroll Admin1

```sh
kubectl exec -n hyperledger deploy/org2-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/users/admin1/msp

kubectl exec -n hyperledger deploy/org2-ca -it -- fabric-ca-client enroll --url https://admin1:admin1pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/users/admin1/msp/

kubectl cp configs/config.yaml hyperledger/fabric-tools:/vol1/organizations/peerOrganizations/org2/users/admin1/msp/
```

### Register Peer0

```sh
kubectl exec -n hyperledger deploy/org2-ca -it -- fabric-ca-client register --id.name peer0 --id.secret peer0pw --id.type peer --url https://admin:adminpw@localhost:7054
```

### Enroll Peer0

```sh
kubectl exec -n hyperledger deploy/org2-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/peers/peer0/msp

kubectl exec -n hyperledger deploy/org2-ca -it -- fabric-ca-client enroll --url https://peer0:peer0pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/peers/peer0/msp/

kubectl cp configs/config.yaml hyperledger/fabric-tools:/vol1/organizations/peerOrganizations/org2/peers/peer0/msp/
```

### TLS Enroll Peer0

```sh
kubectl exec -n hyperledger deploy/org2-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/peers/peer0/tls

kubectl exec -n hyperledger deploy/org2-ca -it -- fabric-ca-client enroll --url https://peer0:peer0pw@localhost:7054 --enrollment.profile tls --csr.hosts localhost --csr.hosts peer0-org2 --mspdir /etc/hyperledger/fabric-ca-client/peers/peer0/tls/
```

### Register Peer1

```sh
kubectl exec -n hyperledger deploy/org2-ca -it -- fabric-ca-client register --id.name peer1 --id.secret peer1pw --id.type peer --url https://admin:adminpw@localhost:7054
```

### Enroll Peer1

```sh
kubectl exec -n hyperledger deploy/org2-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/peers/peer1/msp

kubectl exec -n hyperledger deploy/org2-ca -it -- fabric-ca-client enroll --url https://peer1:peer1pw@localhost:7054 --mspdir /etc/hyperledger/fabric-ca-client/peers/peer1/msp/

kubectl cp configs/config.yaml hyperledger/fabric-tools:/vol1/organizations/peerOrganizations/org2/peers/peer1/msp/
```

### TLS Enroll Peer1

```sh
kubectl exec -n hyperledger deploy/org2-ca -it -- rm -rf /etc/hyperledger/fabric-ca-client/peers/peer1/tls

kubectl exec -n hyperledger deploy/org2-ca -it -- fabric-ca-client enroll --url https://peer1:peer1pw@localhost:7054 --enrollment.profile tls --csr.hosts localhost --csr.hosts peer1-org2 --mspdir /etc/hyperledger/fabric-ca-client/peers/peer1/tls/
```

## Copy `configtx.yaml`

```sh
#kubectl exec -n hyperledger tools -it -- mkdir -p /myvol/orgs/orderers/msp/tlscacerts/
#kubectl exec -n hyperledger tools -it -- mkdir -p /myvol/orgs/orderers/msp/cacerts/
#kubectl cp -n hyperledger config.yaml tools:/myvol/orgs/orderers/msp/
#kubectl exec -n hyperledger tools -it -- cp /myvol/orgs/orderers/users/orderer0/tls/tlscacerts/tls-localhost-7054.pem /myvol/orgs/orderers/msp/tlscacerts/ca.crt
#kubectl exec -n hyperledger tools -it -- cp /myvol/orgs/orderers/users/orderer0/msp/cacerts/localhost-7054.pem /myvol/orgs/orderers/msp/cacerts/
#
#kubectl exec -n hyperledger tools -it -- mkdir -p /myvol/orgs/org1/msp/tlscacerts/
#kubectl exec -n hyperledger tools -it -- mkdir -p /myvol/orgs/org1/msp/cacerts/
#kubectl cp -n hyperledger config.yaml tools:/myvol/orgs/org1/msp/
#kubectl exec -n hyperledger tools -it -- cp /myvol/orgs/org1/users/peer0/tls/tlscacerts/tls-localhost-7054.pem /myvol/orgs/org1/msp/tlscacerts/ca.crt
#kubectl exec -n hyperledger tools -it -- cp /myvol/orgs/org1/users/peer0/msp/cacerts/localhost-7054.pem /myvol/orgs/org1/msp/cacerts/
#
#kubectl exec -n hyperledger tools -it -- mkdir -p /myvol/orgs/org2/msp/tlscacerts/
#kubectl exec -n hyperledger tools -it -- mkdir -p /myvol/orgs/org2/msp/cacerts/
#kubectl cp -n hyperledger config.yaml tools:/myvol/orgs/org2/msp/
#kubectl exec -n hyperledger tools -it -- cp /myvol/orgs/org2/users/peer0/tls/tlscacerts/tls-localhost-7054.pem /myvol/orgs/org2/msp/tlscacerts/ca.crt
#kubectl exec -n hyperledger tools -it -- cp /myvol/orgs/org2/users/peer0/msp/cacerts/localhost-7054.pem /myvol/orgs/org2/msp/cacerts/

kubectl exec -n hyperledger fabric-tools -it -- mkdir -p /vol1/configs/

kubectl cp configs/configtx.yaml hyperledger/fabric-tools:/vol1/configs/
```

## Create Genesis Block

```sh
kubectl exec -n hyperledger fabric-tools -it -- configtxgen -configPath /vol1/configs/ -profile OrdererGenesis -channelID syschannel -outputBlock /vol1/genesis.block
```

## Copy Orderers Configs

```sh
kubectl cp configs/orderer.yaml hyperledger/fabric-tools:/vol1/configs/

#kubectl exec -n hyperledger tools -it -- mkdir -p /myvol/orderers/orderer0/config/
#kubectl cp -n hyperledger config-orderer0.yaml tools:/myvol/orderers/orderer0/config/orderer.yaml
#kubectl cp -n hyperledger configtx-orderer0.yaml tools:/myvol/orderers/orderer0/config/configtx.yaml
#kubectl cp -n hyperledger config.yaml tools:/myvol/orgs/orderers/users/orderer0/msp/
#kubectl cp -n hyperledger config.yaml tools:/myvol/orgs/orderers/msp/
#kubectl exec -n hyperledger tools -it -- cp /myvol/orgs/orderers/users/orderer0/tls/keystore/* /myvol/orgs/orderers/users/orderer0/tls/private.key
```

## Run Orderers

```sh
kubectl apply -f orderer0.yaml
kubectl apply -f orderer1.yaml
kubectl apply -f orderer2.yaml
```

## Create MyChannel

```sh
kubectl exec -n hyperledger tools -it -- configtxgen -configPath /myvol -profile MyChannel -outputCreateChannelTx /myvol/channels/mychannel.tx -channelID mychannel
```