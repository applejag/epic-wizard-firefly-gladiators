///|
struct World {
  mut buttons : @firefly.Buttons
  mut buttons_old : @firefly.Buttons
  mut pad : @firefly.Pad?
  mut pad_old : @firefly.Pad?
  camera : Camera
} derive(Default)

///|
fn World::update(self : World) -> Unit {
  self.buttons_old = self.buttons
  self.buttons = @firefly.read_buttons()
  self.pad_old = self.pad
  self.pad = @firefly.read_pad()
  self.camera.update(self)
}

///|
fn World::render(self : World) -> Unit {
  // Draw background
  @firefly.draw_rect(
    Point::new(0, 0),
    Size::new_screen(),
    Style::new_solid(@firefly.Black),
  )
}
