///|
const SCREEN_WIDTH_HALF_F : Float = 120

///|
const SCREEN_HEIGHT_HALF_F : Float = 80

///|
pub struct Camera {
  pos : Vec2
} derive(Default)

///|
#inline
fn Camera::world_vec2_to_camera_space(self : Camera, position : Vec2) -> Point {
  (position - self.pos).round().to_point()
}

///|
#inline
fn Camera::world_point_to_camera_space(self : Camera, point : Point) -> Point {
  point - self.pos.round().to_point()
}

///|
fn Camera::update(self : Camera, world : World) -> Unit {
  let my_peer = @firefly.get_me()
  let maybe_idx = world.player_system.players_peer.search(my_peer)
  match maybe_idx {
    Some(idx) => world.player_system.players_pos[idx] |> ignore
    _ => ()
  }
  //guard maybe_idx is Some(idx) else { return }
  // Focus on player
  //self.pos = world.player_system.players_pos[idx] //-
  //self.pos = self.pos + Vec2::new(SCREEN_WIDTH_HALF_F, SCREEN_HEIGHT_HALF_F)
  //@firefly.log_debug(
  //  @encoding/utf8.encode(world.player_system.players_pos.to_string()),
  //)
}
