///|
const SCREEN_WIDTH_HALF_F : Float = 120

///|
const SCREEN_HEIGHT_HALF_F : Float = 80

///|
struct Camera {
  mut pos : Vec2
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

}
