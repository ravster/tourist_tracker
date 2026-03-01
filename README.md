# Tourist tracker

Small project to work on my golang & postgres skills. I want to see how much I can bang out with plain old golang stdlib. I usually use Gin for REST & HTTP handling since it is a light library that makes common things easier. But I want to experiment with CQRS for this program, so I'm going to use plain go stdlib. Maybe it'll work out & look good. Or it wouldn't, and I'll have learned something useful.

Nations track the tourist activity in their borders. This is a service that tries to emulate the various functionalities I've seen and heard of as I've travelled.

Hopefully this shows a good example of my Golang and postgres skills. I've done enough Ruby on Rails since 2013 and want to practice more with a statically typed language. Eventually I want to make a script that just shovels a whole load of data at the program, and see how it handles the load and updates aggregate-analysis pages.

# TODO
- Track passport numbers (alnum 20 chars). Age, sex, country, etc to help with profiling visitors.
- Entry port and day (air, sea, land)
- Exit port and day
- Hotel registration by date. Total price paid to hotel. Help identify revenue into the nation's economy by each tourist.
- Police interactions. Because some tourists are annoying. Ban list. Date added.
- Auto-remove a tourist from a ban-list after 5y.
- Event driver (CQRS)
  - I want to experiment with an event-driven system. Writes go to a log, and then other processes write to the actual tables.
  - Each page probably needs to have a refresh button or an auto-refresh for parts of the page so that users see the latest info.
- Analytics
  - Tourists sorted by entry-count over a duration. Identify repeating tourists. Useful for countries like Thailand or Vietnam that have tourists that do visa runs.
  - Sort by duration of stay. This can be useful information to tourism ministries to figure out at a glance if they should modify their visa rules.
  - Sort by revenue. Identify repeating high-revenue tourists and show appreciation for their repeat business. Basic customer appreciation.
  - Highest revenue for duration by country. Might want to do more advertising to get more of these tourists.
  - Higest offenders (police interactions, ban list) by country. Might want to keep some people out.

# DONE
