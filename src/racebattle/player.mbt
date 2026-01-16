///|
/// Rotation speed. Angle is the rotation speed per second
let rotation : Angle = Angle::from_deg(180.0 / 60.0)

///|
const MOVE_SPEED : Float = 5

///|
pub struct Player {
  // manually doing "struct of arrays" optimization,
  // mostly as an experiment to see how tedious it is,
  // but also because MoonBit doesn't allow "#valtype"s in other "#valtype"s
  peer : @firefly.Peer?
  mut pos : Vec2
  mut angle : Angle
}

///|
pub fn Player::new(peer : @firefly.Peer?) -> Player {
  Player::{ peer, pos: Vec2::default(), angle: Angle::default() }
}

///|
pub fn Player::update(self : Player) -> Unit {
  @firefly.log_debug(@encoding/utf8.encode("\{self.peer}: player update"))
  guard self.peer is Some(peer) else { return }
  @firefly.log_debug("peer is some")
  let pad_maybe = @firefly.read_pad(peer~)
  guard pad_maybe is Some(pad) else { return }
  @firefly.log_debug("pad is some")
  guard pad.x != 0 && pad.y != 0 else { return }
  @firefly.log_debug("pad not zero")
  let new_angle = @util.rotate_toward(self.angle, pad.azimuth(), rotation)
  let dir = Vec2::from_angle(new_angle).scale(MOVE_SPEED)
  self.pos = self.pos + dir
  self.angle = new_angle
  @firefly.log_debug(@encoding/utf8.encode("\{peer}: Pad(\{pad})"))
}

///|
pub fn Player::render(self : Player, world : World) -> Unit {
  let point = world.camera.world_vec2_to_camera_space(self.pos)
  let style = if Some(world.me) == self.peer {
    @firefly.Style::new_solid(DarkGreen)
  } else {
    @firefly.Style::new_solid(Red)
  }
  @firefly.draw_circle(point, 3, style)
}
