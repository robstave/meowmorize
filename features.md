
Need
------------------
Dark mode, the cat graph has no ears.  Darkmode needs a white line, not black

Edit a card.  In the dialog and message, a flag to reset the card stats

sparkline to show the total stats of a deck. Average I guess.

Create new deck says default.  it needs a dialog for a name

deck updated timestamp

add date ( above) to decks

add bookmark to deck

add stars to the bar
https://jsfiddle.net/sqvgLw98/1/


new concept  ( notes )

add name of deck to card page or whatever

click in to rename

https://github.com/remarkjs/react-markdown#use

remove DecksDashboard?

consider adding a title to the markdown
<!-- title:foo -->

logging session stats
This is there, but kinda coupled and does not really work for resort.
maybe do all that processing in a channel.


ByteBeat concept notes
------------------
- Rename/brand concept as **ByteBeat**.
- Core audio texture idea: bake ByteBeat formulas into gameplay.
- Candidate formula (8 kHz):
  `t*(t^t+(t>>15|1)^(t-1280^t)>>10)`
- Enemy idea: add a sphere enemy that emits ByteBeat audio.
  - Volume increases as the sphere gets closer to the player.
  - The enemy "texture" / vibe is represented by the generated waveform.
- References:
  - https://medium.com/small-tech/bytebeat-experiments-9d11eee630c8
  - https://zopium.neocities.org/bytebeat/?KHQ+PjZ8dHx0Pj4odD4+MTYpKSoxMCsoKHQ+PjExKSY3KSx8fHx8fDgwMDAsfHx8fHxi


----
progress

func (r *SessionLogRepositorySQLite) GetSessionLogsBySessionID(sessionID string) ([]types.SessionLog, error) {
    var logs []types.SessionLog
    err := r.db.Where("session_id = ?", sessionID).Order("created_at ASC").Find(&logs).Error
    if err != nil {
        return nil, err
    }
    return logs, nil
}

func (r *SessionLogRepositorySQLite) GetSessionLogsByUser(userID, deckID string) ([]string, error) {
    var sessionIDs []string
    


    ----

done
-----------------
added clone, update and delete apis

collapse decks

