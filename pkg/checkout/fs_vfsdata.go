// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package checkout

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// FS statically implements the virtual filesystem provided to vfsgen.
var FS = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/css": &vfsgen۰DirInfo{
			name:    "css",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/css/global.css": &vfsgen۰CompressedFileInfo{
			name:             "global.css",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 2743,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x55\xdd\x8a\xe3\x38\x13\xbd\xf7\x53\x88\x19\x3e\xe8\x34\x56\xda\x71\x7e\x3a\xed\xef\x66\xe9\x85\x81\x85\x6d\x58\x66\x60\x6f\x87\xb2\x5d\x8e\xb5\x91\x25\x23\x29\xe9\x78\x96\xbc\xfb\xa2\x1f\xc7\x76\x92\x59\x1a\x16\x9a\xb4\x5d\x3a\xb2\x4e\x1d\x55\x9d\x7a\x7a\x24\x7f\x82\x62\x90\x73\xd4\xe4\xf1\x29\xca\x94\x94\x86\xfc\x1d\x11\x42\xe9\x4e\x41\x47\x65\x55\x69\x34\x19\x51\xbb\x1c\x1e\x92\x98\x84\xbf\x79\xb2\x9c\xfd\xdf\xa1\x72\x59\x76\xb4\x90\x5c\x2a\x07\x7a\x58\xbf\xc4\x64\xb3\x8e\xc9\xcb\x2a\x00\x6a\x84\x92\x33\x81\x63\xd0\x72\x11\x93\x74\x13\x93\xe7\x45\x00\x41\x51\xa0\x30\x3d\xe4\x73\x92\x6c\x36\x55\x32\x3a\xa0\x92\xc2\xd0\x0a\x1a\xc6\xbb\x8c\x50\x68\x5b\x8e\x54\x77\xda\x60\x13\x93\x57\xce\xc4\xfe\x0d\x8a\x6f\xee\xfd\x8b\x14\x26\x26\x1a\x84\xa6\x1a\x15\xab\xfc\x47\x14\x94\xec\xa0\x33\xb2\x69\x4f\x3e\x50\x49\xd5\xd0\x77\x56\x9a\x3a\x23\x9b\x24\xb1\xe1\x73\x14\x3d\x3d\x92\x57\xd0\x68\x95\x78\x74\x2a\xe4\xf2\x44\x35\xfb\xc1\xc4\x2e\x23\xb9\x54\x25\x2a\x9a\x4b\x87\xb5\xb4\x1c\x64\x42\xed\x08\xea\xe1\x96\xb2\x4b\xd2\xbd\x6b\xf6\x03\x33\xb2\x08\x34\x42\xba\xa3\x4d\x2e\xe2\x35\x79\xc7\x7c\xcf\x8c\xff\x8c\x6e\xa4\x34\xb5\x63\x01\xc2\x30\xe0\x0c\x34\x96\x96\x46\xbd\x88\xa3\x3a\x8d\xa3\x7a\x19\x47\xf5\x2a\x8e\xea\x75\x1c\xd5\x1b\x47\x6c\xf2\xf5\xe9\x2d\x0c\x84\xde\x91\xed\x6a\x93\x11\x21\x55\x03\xdc\x86\x1b\x50\x3b\x26\xa8\x91\x6d\x46\x52\xcf\x33\x84\x72\x69\x8c\x6c\x32\xb2\xf2\x6a\xd5\x8b\x21\x7f\x9f\xd7\x32\x0d\x2b\xe9\xf5\x4a\x1a\xf6\x44\x07\x3e\xe6\xf6\x79\xfb\xfc\x92\x42\x3a\x9c\x91\x11\x77\xeb\x2d\x94\xa5\xcb\x36\x09\x9b\xe6\x42\xd2\xfc\xc0\x39\x1a\xed\xf6\x73\xa6\x0d\xd5\xa6\xe3\x48\x4d\xd7\xa2\xe5\x2f\xb0\x3f\x80\xb3\x80\x11\x48\xeb\x90\xdf\x98\x00\x67\xa4\x64\xc7\x5b\x8d\x46\x37\x10\x90\x57\x6c\x8b\x45\xf1\x52\x3a\xb6\x37\xe7\x97\x4c\x17\x23\xa9\x38\x56\x26\x23\x8b\xed\x50\x57\xbf\x43\x27\x0f\xc6\x56\xd6\x5c\x2b\x7a\x69\xb3\x92\xe9\x96\x43\x97\x91\x8a\xa3\xd3\xda\xfe\xa7\x25\x53\x58\x18\x26\x45\x46\x94\x7c\xb7\xe1\x50\xaa\x8b\x24\xf9\x9f\x3f\xe6\xd4\x57\xef\xcb\x36\xf1\xb7\x74\x11\x6d\xb5\xf5\x01\xe0\x6c\x27\x68\x21\x85\x41\x61\x32\x62\x3b\x0c\x95\x5d\xf8\xeb\xa0\x0d\xab\xba\x7b\x4b\xbd\x5e\x70\x30\xd2\x1d\xc4\xc4\x45\xc3\x45\x92\x1c\xeb\xc9\x5d\x05\xd8\xd9\xe5\x64\x4b\x0c\x95\xcb\xea\xaa\x60\xfa\xba\xb0\xa8\x16\xba\xc6\xb6\xba\x3e\x34\x0d\xa8\xee\x1e\x3c\x4d\x06\x78\x03\x4c\xc4\xee\x29\x90\xfd\x98\x6a\x85\xe4\x87\x46\x7c\x30\xd7\x5e\x54\xaf\x97\x46\x5e\x0d\xa8\x81\x85\x3b\x38\x68\xee\xeb\x65\xf0\x90\xd9\x54\x83\xef\xdf\xb9\xdc\x49\x6f\x20\x50\xec\x77\x4a\x1e\x44\x49\x59\x03\x3b\xec\xf7\x5a\x80\x8f\xcc\xc6\x54\x7c\x99\x4e\xb6\xf9\x06\xb2\xf4\x81\x89\xab\x35\x85\x2d\x82\xeb\xde\xf0\x78\x53\x2a\xa1\xf8\x98\xd8\x3b\x7f\x87\xdb\xaa\x1f\x5b\xaf\xe3\x62\xf0\x64\x68\x89\x85\x54\xe0\xc5\xf4\xbd\x45\x88\x51\x20\x34\xf3\x31\xe0\x9c\x24\xf3\x54\x13\x04\xed\x1b\x0f\xb2\x5a\x1e\xc3\xfd\x57\x8c\x1b\x54\x19\xc9\x95\x4d\x4b\xa0\xd6\x0f\xc9\x7c\x3b\x0b\x38\x28\x0c\x3b\xe2\xcf\x81\xeb\x59\xcf\xfb\x57\x59\x22\xc9\xb9\x2c\xf6\x7d\xe3\x14\xc0\xb9\xed\xa3\xa9\xb6\x7d\x2e\xa3\x91\x35\x9b\xb4\xc4\x22\x38\x59\xf0\xf0\x7e\x1c\xf8\x5d\xfe\x6d\xd6\xf7\xd5\xe5\x2e\x92\xd0\x58\x36\xad\x8a\xcb\xf7\xec\x52\xee\x51\x21\x4b\x8c\xa3\x56\x61\x1c\xe9\x16\xc4\xdc\xbe\xdf\x70\xb2\x93\x2e\x5d\x3d\xc7\x24\x5d\x27\xf6\x27\x9d\x7d\x84\x82\x5f\xcf\xc8\xa2\x3d\x11\x2d\x39\x2b\xfd\x77\x52\xfb\x9d\x65\x6a\x7f\xb6\x83\x7d\xf7\x73\xe7\xd3\xb7\x2f\xe4\x4d\x0a\xf9\x29\x26\x9f\x7e\x7b\x7d\x23\x7f\x70\x3c\x5d\x02\x6f\x28\xb8\x7d\x68\xa4\x90\xba\x85\x02\xa7\x29\x38\xde\x17\xa5\x36\xed\x29\xa8\x75\x8e\xae\x72\x1b\x0c\x26\x2c\xdb\x95\x79\x09\x6a\xef\xbe\xe3\x9e\x6e\x44\xf8\xbc\xaa\xd6\x9b\x4d\x3e\x1a\x77\x2e\x9d\xf5\xda\x2a\xe2\x7f\x6e\xd3\x71\x8c\xe9\x57\xdc\x1d\x38\xa8\xeb\xe1\xb9\xf4\xb7\xf2\x93\xf1\xc5\xd1\x18\x54\xd4\xe6\xe9\x27\x88\x47\x4f\x87\x41\x72\x37\xbf\xff\xaa\xe8\xd3\x23\xf9\x8a\xba\x95\x42\xb3\x23\xda\x5a\xb6\x45\xfb\x4b\x83\x25\x03\xf2\x30\x72\xec\x67\x7b\xfe\xcc\x1d\x39\x99\x05\xff\x66\x63\x77\x8c\xcc\x81\xb5\x01\x65\x3c\x60\xe2\xff\x21\x47\x1b\xb7\x16\x1e\x4e\x5e\xf6\xd1\x73\x14\x0e\xbf\x31\xac\x89\xbf\xb4\xb2\xef\xf7\xc1\x35\x2f\x5b\xef\x39\x79\x70\x0f\xe7\xa4\x77\x37\x8d\x5d\x7c\xe4\xe3\xbd\xc5\x5c\x70\x17\xc7\xbd\x19\x7d\xe7\xe8\x1c\xfd\x13\x00\x00\xff\xff\x3e\x45\x79\xc6\xb7\x0a\x00\x00"),
		},
		"/html": &vfsgen۰DirInfo{
			name:    "html",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/html/cancel.html": &vfsgen۰CompressedFileInfo{
			name:             "cancel.html",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 910,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\x4b\x4f\xe3\x3e\x10\xbf\xf3\x29\xe6\xef\x7b\x33\x7f\xb4\x1c\x10\x72\x22\xad\x60\x0f\x9c\x76\x25\xf6\xc2\x09\xb9\xf6\xb4\x71\xf1\x23\xf2\x4c\x0b\x5d\xc4\x77\x5f\xa5\x4e\x69\x0a\x97\x3d\xc5\x9e\xdf\x63\x5e\x8e\xfe\xef\xee\xe7\xed\xef\xc7\x5f\x3f\xa0\x97\x18\xba\x0b\x3d\x7e\x20\x98\xb4\x6e\x15\x25\xd5\x5d\x00\xe8\x9e\x8c\x1b\x0f\x00\x3a\x92\x18\xb0\xbd\x29\x4c\xd2\xaa\xad\xac\x16\xd7\x0a\x70\x02\xc5\x4b\xa0\xee\x41\x8a\x1f\x08\x6e\x7b\xb2\xcf\x79\x2b\x1a\x6b\x78\xa6\x4f\x26\x52\xab\x1c\xb1\x2d\x7e\x10\x9f\x93\x02\x9b\x93\x50\x92\x56\x7d\x07\x47\x31\x43\x5e\xc1\x27\x9f\x43\x9a\x6a\x12\x7c\x7a\x86\x42\xa1\x55\xde\x8e\xe2\xbe\xd0\xaa\x55\x2b\xb3\x1b\xaf\x8d\xb7\x59\x81\xec\x07\x6a\x95\x8f\x66\x4d\xf8\xba\xa8\xb4\x63\x99\x27\x39\xcb\x3e\x10\xf7\x44\x72\x34\xe9\x45\x06\xbe\x41\xb4\x2e\x6d\xb8\xb1\x21\x6f\xdd\x2a\x98\x42\x8d\xcd\x11\xcd\xc6\xbc\x62\xf0\x4b\xc6\x94\x4b\x34\xc1\xff\x21\xbc\x6e\xfe\x6f\x2e\x4f\xf7\x26\xfa\xd4\x58\xe6\x7f\xcb\x86\x2c\x46\xbc\x45\xcb\x8c\xeb\x90\x97\x26\x9c\x6b\xeb\x80\x80\x8b\x3d\x55\xb6\xe1\x86\x0f\x93\x39\x94\xb4\xfb\x86\xaa\xd3\x58\x89\x87\x65\x61\xdd\xd6\x78\x5c\x66\xb7\x9f\x9c\x9c\xdf\x81\x0d\x86\xb9\x55\x5c\x16\x25\x67\x51\x15\xf9\x82\x45\xe3\xd3\x07\x36\x2d\x9f\xca\x8c\x50\x03\x33\xca\x17\x8b\xca\x78\x7a\x0a\x79\x9d\xc7\xea\x9c\xdf\xcd\x0c\xb1\xc2\xb3\xc8\xb9\x7a\x30\xfb\x48\x49\x16\xbc\x8d\xd1\x94\x3d\xd8\x1c\x87\x40\x42\x6e\xb1\xf3\xf4\x72\x9e\xb7\xbf\xec\x1e\xf3\xb6\x80\x10\x0b\x4c\x42\x78\x31\x0c\xd6\x24\x4b\x81\x9c\xc6\xfe\xf2\x5c\x71\xd5\x1d\x5f\x14\x3c\x10\xb3\xcf\x09\xee\xef\x6e\x40\xf3\x60\x12\x78\xd7\x2a\xae\x51\xd5\xbd\xbd\x35\x13\xe3\xfe\xee\xfd\x5d\xe3\xc8\xe8\x34\xf6\x57\xf3\x66\x66\xbd\x7d\xea\x63\x7a\xd2\xa7\x39\x9f\xb8\x1f\x47\x8d\x75\x47\x1a\xeb\xdf\xf7\x37\x00\x00\xff\xff\xf3\x28\x79\x86\x8e\x03\x00\x00"),
		},
		"/html/redirect.html": &vfsgen۰CompressedFileInfo{
			name:             "redirect.html",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 482,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x91\xb1\x6e\xc2\x30\x10\x86\x77\x9e\xe2\xea\x85\x30\x34\x1e\xba\x54\xe0\x44\xaa\xa0\x43\xd5\xa1\x48\xb0\x74\x34\xce\x81\xdd\x3a\x36\xb2\x2f\x48\x08\xe5\xdd\xab\x60\xa7\xd0\x4e\x77\xe7\xfb\xbf\xff\x7c\x3a\xf1\xb0\xfa\x58\x6e\x3f\xd7\xaf\xa0\xa9\xb5\xf5\x44\x0c\x01\xac\x74\x87\x8a\xa1\x63\xf5\x04\x40\x68\x94\xcd\x90\x00\x88\x16\x49\x82\xd2\x32\x44\xa4\x8a\x75\xb4\x7f\x7c\x66\xc0\x73\x93\x0c\x59\xac\x37\x14\xcc\x11\x61\xa9\x51\x7d\xfb\x8e\x04\x4f\xcf\x77\xbc\x93\x2d\x56\xac\xc1\xa8\x82\x39\x92\xf1\x8e\x81\xf2\x8e\xd0\x51\xc5\x5e\xa0\xc1\xd6\x83\xdf\xc3\x3f\x9f\xdb\x98\x84\x41\x0c\xaa\x62\x9a\xe8\x18\xe7\x9c\x7f\xc5\x32\x5e\xf5\xa5\xf2\x2d\x3f\x3d\x71\x56\x0b\x9e\x84\x7f\xa8\x54\x00\x9c\x64\x80\x04\x40\x95\x27\x15\xd3\xcb\xa5\x5c\x77\x3b\x6b\xa2\x96\x3b\x8b\xef\x78\xee\xfb\xe9\x6c\x91\x89\x6c\x1f\xb0\x31\x01\x15\x6d\xfd\xf8\xb1\xe2\x92\x15\x00\x11\x63\x34\xde\xbd\x35\x73\x18\xcc\x36\xb9\x5c\xf5\xfd\x34\x6b\xfa\x59\x49\x1a\x5d\xb1\xef\x9c\x1a\x36\x87\x22\x60\xec\x2c\xcd\xe0\xe6\xa2\xbc\x8b\xde\x62\x69\xfd\x61\xec\x2e\x7e\xf1\x94\xdd\xef\x26\xf8\x78\x1f\xb1\xf3\xcd\xb9\x16\xfc\x1a\x26\x82\xa7\x8b\xfe\x04\x00\x00\xff\xff\xb2\x15\xd3\xdd\xe2\x01\x00\x00"),
		},
		"/html/success.html": &vfsgen۰CompressedFileInfo{
			name:             "success.html",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 2873,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\xd1\x6f\xdb\xb6\x13\x7e\xef\x5f\x71\x25\x7e\xaf\x12\x93\xf4\xa5\x08\x24\xfd\x30\x24\xc3\x10\xa0\xdd\x8a\xb5\x7d\x28\xb2\x20\xa0\xc9\xb3\xc5\x98\x22\x35\x1e\xe5\xc4\xeb\xfa\xbf\x0f\x14\x25\x5b\x96\xe3\x2e\x28\xf6\x60\x58\xd4\x7d\x77\xf7\xdd\xf1\xe3\x51\xc5\xeb\xeb\xdf\xae\x3e\x7d\xf9\xf0\x33\xd4\xa1\x31\xd5\xab\x22\xfe\x81\x11\x76\x55\x32\xb4\xac\x7a\x05\x50\xd4\x28\x54\x7c\x00\x28\x1a\x0c\x02\x64\x2d\x3c\x61\x28\x59\x17\x96\xd9\x5b\x06\x7c\x30\x06\x1d\x0c\x56\x1f\x83\xd7\x2d\xc2\x55\x8d\x72\xed\xba\x50\xf0\xf4\x7a\xe2\x6f\x45\x83\x25\x53\x48\xd2\xeb\x36\x68\x67\x19\x48\x67\x03\xda\x50\xb2\x9f\x40\x61\xe3\xc0\x2d\x61\x16\xa7\x4f\x93\x82\x18\x6d\xd7\xe0\xd1\x94\x4c\xcb\xe8\x5c\x7b\x5c\x96\x6c\x29\x36\x71\x99\x6b\xe9\x18\x84\x6d\x8b\x25\xd3\x8d\x58\x21\x7f\xca\x12\x6c\xa4\xb9\x77\xa7\xb0\x35\x48\x35\x62\x18\x83\xd4\x21\xb4\x74\xc9\xb9\x54\xf6\x81\x72\x69\x5c\xa7\x96\x46\x78\xcc\xa5\x6b\xb8\x78\x10\x4f\xdc\xe8\x05\x71\xeb\x7c\x23\x8c\xfe\x0b\xf9\xdb\xfc\x2c\x3f\xdf\xaf\xf3\x46\xdb\x5c\x12\xed\xb3\xbd\xce\xb2\xef\x66\xe4\x14\x44\xd0\x92\x4b\x22\xbe\x32\x6e\x21\xcc\xe8\x9f\x65\x2f\x20\xbc\xd4\x06\x2f\x39\xe7\x9f\x09\x3d\xf1\xe0\x1a\xf4\x5c\x3a\x85\x9c\xfa\xf6\x65\xd2\x68\xde\xae\x57\x5c\x0e\x6d\x3c\x9d\x2e\xe5\x4a\x7b\x02\xe4\xe5\xbe\x19\x0f\x94\xa7\x68\x7d\x17\x36\x6f\x38\xab\x0a\x9e\x80\xbd\x3e\x78\x12\x48\x7c\x5c\x38\xb5\x1d\x22\x29\xbd\x01\x69\x04\x51\xc9\xc8\x67\xde\xb9\xc0\x92\xe5\xc8\xd6\x08\x6d\x77\xb6\x41\x6f\xe8\x27\x80\xf4\x62\x02\x39\x0a\x91\x10\xf7\xf7\xc6\xad\x5c\x64\xa7\xf4\x66\x12\x90\x27\xf3\xe4\xcd\xa1\x77\x2b\xb6\x0d\xda\x90\x51\xd7\x34\xc2\x6f\x41\xba\xa6\x35\x18\x50\x65\x1b\x8d\x8f\x87\x79\xeb\xf3\xea\x8b\xeb\x3c\x04\xa4\x00\x83\x23\x50\x27\x25\xa2\x42\x55\xf0\xfa\x7c\x4e\x73\xba\x06\x28\xda\x6a\x94\x34\x7c\x44\x22\xed\x2c\xdc\x5c\x5f\x16\xbc\x9d\xe3\x3c\x56\x5f\xbf\xe6\x03\xe6\xe6\xfa\xdb\xb7\x82\xc7\x77\xd3\xe8\xf3\x3a\x27\xcb\x59\x89\xc3\x01\x3b\x68\xf3\x45\xf5\x2b\x3e\x05\xa0\x80\x2d\x15\xbc\xbe\x98\xd8\x3a\x73\x90\xc7\xe8\x19\xb9\x67\xcb\xf2\x28\x02\x82\x00\xd9\x51\x70\x0d\x8c\x92\x03\x1a\xca\x2c\xc4\xec\x94\x4d\x54\xa5\x9c\x24\x2e\x5a\x3d\x11\x6a\xf2\x22\x2e\xfb\xb8\xff\xef\x67\x92\x34\x9a\x55\x8f\x3a\xd4\x10\x6a\x84\xab\x77\x37\x97\x05\x17\xd5\xb3\xcd\x1b\x8b\x57\xc2\xaf\x59\xf5\x3f\x48\xc9\x8e\x58\x11\xa4\xf8\xf0\x47\x1f\x21\xcb\xfa\xbd\x24\xca\x3a\x6f\xf6\x4c\xf1\x49\x44\x49\xf4\x54\x07\x00\xdb\x79\x48\x61\x25\x9a\xd3\x0e\xc9\xbe\xc3\x2b\x60\x83\x6e\xee\x1b\x0c\xb5\x53\xf7\x71\x5a\xd1\xed\x1d\x2b\xa5\xf0\x6a\x02\x33\xda\xe2\xbd\x0e\xd8\xd0\xed\xd9\xdd\x6d\x1c\x9b\x77\xac\x64\xc3\x60\xfc\xe0\xa7\x24\x8e\xd0\x93\xf1\x1a\x9d\x7e\xc1\x00\x9f\xdc\x0a\x43\x8d\xfe\xb4\x93\x68\x5c\x67\xc3\x1d\x2b\x2f\xde\x9c\x9d\x9d\x44\xc9\xce\x7b\xb4\x72\x7b\xc7\xca\x8e\x4e\xf3\xfd\xb3\x13\x36\xe8\x10\x61\x17\x3d\xe8\xef\x61\x0f\x32\x85\x9b\xfd\x3e\xf8\xce\x0e\xc2\x9e\xc9\x39\x0a\xfa\x50\x78\x2f\xd4\xa1\x71\x16\xa1\xff\x2d\x21\x1e\xd5\xfd\x96\xf7\x7b\x42\x30\x15\x10\x04\x07\x84\x08\x8d\xf3\x08\xda\x2a\x6c\x43\x0d\x1d\x21\x48\x41\x48\xc7\xca\x3a\x3c\x19\x23\xab\x23\x69\xaf\x74\xa8\xbb\x45\xd2\x4b\x2a\x7a\x48\xbe\x13\x78\xe6\x2c\x66\x41\x37\x38\x8e\x1f\x62\xd5\x69\x5b\x92\xf9\xbc\xfe\x1f\x4f\x4e\xda\xae\x0c\x66\xd4\x2d\xf6\xb7\x70\xf5\x3d\xeb\x7f\x4e\x60\x12\x3b\x13\x56\x65\x42\xa9\xec\x90\xc4\xf3\x88\xe7\x89\x14\x7c\x36\xb2\x5e\xa2\xa5\x5e\x3e\x45\x5b\xbd\x17\x6b\x04\xea\x3c\x46\x31\x48\x67\x97\x7a\x15\x17\x02\x1e\x71\x51\x3b\xb7\x26\x40\xab\x5a\xa7\x6d\x80\xa5\xf3\x50\x50\x2b\xec\x38\x5e\xe2\x65\xbb\x27\x9d\x0f\x53\x25\xdf\x5d\x21\x05\x8f\xe8\x0a\x84\x55\xbd\xac\x42\x2d\x42\x4c\x23\xf6\xb7\xcc\x78\x8d\xe4\x51\x6d\x89\xf8\x21\xdb\x54\xdc\xc1\xe4\x7e\x1f\xe5\xea\x51\x28\x6d\x57\x47\xb3\x7b\xe4\x66\x5d\xb6\xe8\x8c\xc1\x28\xad\x79\xe9\xff\x36\x8b\x47\xe1\xed\x76\x8c\xcd\xbf\xec\x40\x39\xd9\x45\x8c\x38\xa5\x8f\x17\x25\x1a\x7b\xcc\xaa\xdf\x51\xa2\xde\x68\xbb\xda\xf5\xfd\x87\x83\xee\xbf\x7e\x58\xf5\x39\xaa\xb9\x3f\xee\x63\x01\xef\x6e\x8e\x03\x4f\x05\x34\x11\xcf\xee\xb1\xe0\xe9\xdb\xa6\xe0\xe9\x43\xf9\x9f\x00\x00\x00\xff\xff\x25\xd1\x1f\x89\x39\x0b\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/css"].(os.FileInfo),
		fs["/html"].(os.FileInfo),
	}
	fs["/css"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/css/global.css"].(os.FileInfo),
	}
	fs["/html"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/html/cancel.html"].(os.FileInfo),
		fs["/html/redirect.html"].(os.FileInfo),
		fs["/html/success.html"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}