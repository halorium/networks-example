package serializers

import "io"

type Deserializer func(io.ReadCloser, interface{}) error

type Serializer func(interface{}) []byte
