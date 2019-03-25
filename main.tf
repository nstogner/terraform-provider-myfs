resource "myfs_text_file" "my-file" {
  path = "/tmp/myfile.txt"
  text = "Hi!"
}
