A simple tool that calculates the fastest and earliest connection between two locations in Switzerland, by using public transport.

Usage:
  connections [OPTIONS]

Application Options:
      --from=     declare the name of the departing city
      --to=       declare the name of the destination city
      --time=     declare time of departure, format [hh:mm]
      --date=     declare date of departure, format [YYYY-MM-DD]
      --arriveby  optional, mark declared time and date as the arrival ones, usage
                  --arriveby
      --direct    optional, query only direct connections, usage --direct

Help Options:
  -h, --help      Show this help message

#EXAMPLES
1. Find connections from Geneva to Lausanne in the defined time and date

go run connections.go --from Geneva --to Lausanne --time 13:00 --date 2021-06-23

2. Find connections from Geneva to Lausanne in the defined time and date, but consider only the direct ones

go run connections.go --from Geneva --to Lausanne --time 13:00 --date 2021-06-23 --direct

3. Time and date are the arrival ones

 go run connections.go --from Geneva --to Lausanne --time 13:00 --date 2021-06-23 --arriveby

 

Output:
******FASTEST CONNECTION******
FROM: Genève
TO: Lausanne
DEPARTURE TIME: 2021-06-23T13:39:00+0200
DURATION: 00d00:36:00
TRANSFERS: 0
DELAY: 0
DEPARTURE PLATFORM: 6
ARRIVAL PLATFORM: 6

******EARLIEST CONNECTION******
FROM: Genève
TO: Lausanne
DEPARTURE TIME: 2021-06-23T13:11:00+0200
DURATION: 00d00:37:00
TRANSFERS: 0
DELAY: 0
DEPARTURE PLATFORM: 6
ARRIVAL PLATFORM: 6

#WHAT I WOULD DO DIFFERENTLY
1. Add a graphical interface instead of CLI.
2. Take into consideration also other params, such as delay, transfers etc.
