///|
const FPS : Float = 60

///|
/// Rotation speed. Angle is the rotation speed per second
let rotation : Angle = Angle::from_deg(120.0 / FPS)

///|
const MOVE_SPEED : Float = 0.8

///|
pub struct Player {
  // manually doing "struct of arrays" optimization,
  // mostly as an experiment to see how tedious it is,
  // but also because MoonBit doesn't allow "#valtype"s in other "#valtype"s
  sprite_sheet : @util.SpriteSheet
  sprite_sheet_rev : @util.SpriteSheet
  peer : @firefly.Peer?
  mut pos : Vec2
  mut angle : Angle
}

///|
pub fn Player::new(
  sprite_sheet : @util.SpriteSheet,
  sprite_sheet_rev : @util.SpriteSheet,
  peer : @firefly.Peer?,
  pos? : Vec2 = Vec2::default(),
) -> Player {
  Player::{ sprite_sheet, sprite_sheet_rev, peer, pos, angle: Angle::default() }
}

///|
pub fn Player::update(self : Player) -> Unit {
  self.sprite_sheet.update()
  self.sprite_sheet_rev.update()
  if self.peer is Some(peer) {
    let pad_maybe = @firefly.read_pad(peer~)
    if pad_maybe is Some(pad) {
      if pad.x != 0 && pad.y != 0 {
        let new_angle = @util.rotate_toward(self.angle, pad.azimuth(), rotation)
        self.angle = new_angle
      }
    }
  }
  let dir = Vec2::from_angle(self.angle).scale(MOVE_SPEED)
  self.pos = self.pos + dir.scale(MOVE_SPEED)
}

///|
pub fn Player::render(self : Player, world : World) -> Unit {
  let point = world.camera.world_vec2_to_camera_space(self.pos)
  // draw shadow
  @firefly.draw_circle(
    point + Point::new(-2, 2),
    5,
    @firefly.Style::new_solid(DarkGray),
  )
  // draw arrow of movement direction
  let dir = Vec2::from_angle(self.angle)
  @firefly.draw_line(
    point + dir.scale(6).to_point(),
    point + dir.scale(9).to_point(),
    @firefly.LineStyle::new(Black),
  )
  // draw sprite
  if Some(world.me) == self.peer {
    let is_looking_left = Angle::from_rad(@micromath.PI)
      .difference(self.angle)
      .to_rad()
      .abs() <
      @micromath.FRAC_PI_2
    if is_looking_left {
      self.sprite_sheet_rev.draw(point - Point::new(4, 5))
    } else {
      self.sprite_sheet.draw(point - Point::new(4, 5))
    }
  } else {
    @firefly.draw_circle(
      point - Point::new(2, 2),
      5,
      @firefly.Style::new_solid(Red),
    )
  }
}
