Sử dụng FlatBuffers (https://google.github.io/flatbuffers/) để serialize một số dữ liệu.
- https://github.com/secmask/contact
    một file ".idl" là mô tả format của message trong flatbuffers,
    một file "*.go" chứa type tương ứng trong Go.

- viết một lib để:
1. serialize object là Go type ở trên cho dùng flatbuffers,
2. de-serialize data đã output ở yêu cầu số 1 sang Go type ban đầu.

Tức là convert 2 chiều.
