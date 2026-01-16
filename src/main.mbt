///|
let racebattle_world : Ref[@racebattle.World] = Ref::new(
  @racebattle.World::default(),
)

///|
fn main {

}

///|
pub fn boot() -> Unit {
  //let assets = Assets::load()
  racebattle_world.val = @racebattle.World::{
    ..racebattle_world.val,
    player_system: @racebattle.PlayerSystem::new(@firefly.get_peers()),
  }
}

///|
pub fn update() -> Unit {
  @firefly.log_debug("update start")
  guard racebattle_world.val.player_system.players_peer.length() > 0 else {
    @firefly.log_error("crap")
    return
  }
  @firefly.log_debug("before world update")
  racebattle_world.val.update()
  @firefly.log_debug("update end")
}

///|
pub fn render() -> Unit {

}
