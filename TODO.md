# Swiss 1.1.0 and 1.0.3

## Misc:
Github: \
    1. Create projects on issue board 


1.0.3: \
    1. Consolidate modules into utils ( colors, messages, helps maybe ). - DONE
    


1.1.0: \
    1. "Net" module: \
        1.1: Port scanner \
        1.2: IP Address Pinging \
        1.3: WHOIS Lookup \
        1.4: HTTP 'Probe': Similar to a curl request to grab http status code and http information. \
        1.5: Advanced Ping: Similar to the default systems ping but provides more information regarding to packets and other related network information. \
    2. "Generator" module \
        2.1: Passwords \
        2.2: UUID \
        2.3: Hex Codes ( Based on main color input or random otherwise ) \
        2.4: "Secret" codes, strings of letters, numbers and symbols created based on users number. 



Command struct
    -- Name string
    -- Subcommands []Subcommand


Subcommand struct 
    -- Name ( flag or argument )
    -- Handler func(x y or z)
