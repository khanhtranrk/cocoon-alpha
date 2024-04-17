# cocoon-alpha
cocoon alpha

# Nw
Cocoon
    Gates
    Taistra
    Justice
    ...

# Tb
Cocoon
    Citizen (Taistra, Justice)
    Contact (All)
    Letter
        Unprocessed Letter
        Unsent Letter
        Processed Letter
        Sent Letter
        Refer Letter

# St
Cocoon
    Broker {
        AmqpUrl
    }

    ContactGate {
        Broker
        SendMessage(msg, gate)
    }

    Citizen {
        Id
        Name
        ContactGate
        SerectKey
        JoinTime
    }

## Domains

### Citizen

```
Id          "TASKA"
Name        string
ContactGate uint64
ActiveDate  datetime
```

### Contact

```
CitizenId
ContactGate
```

taska
    endcodeFemix
    decodeFemix

taska.BuildLetter(msg) {
    ct, err := cocoon_service.GetContactByName("TASKA")
    ct.Id
}

femix nartual -> byte
femix nartual
if femix nartual to byte
dafdsfdsafsa
    if sadfsadfsad
        taska.femixnartual to byte ...


Contact
Id
Key -> verify message from contact

Cocoon.ConnectToCocoon() -> tcp
