### Save Video Api
This API provides a solution for streaming video data from a client-side application, collecting video chunks at regular intervals, merging those chunks into a complete WebM video file, saving the final file to the server's disk, and then transcribing the video content using OpenAI's Whisper API with the help of RabbitMQ for task queuing.

#### Endpoints
##### Start Stream
Request
```
curl http://localhost:8080/api/startStream
```
Response
```json
{
    "data": {
        "video_id": "vid_57c2cca2-0908-43b9-87ac-fe2229d1aec7"
    },
    "message": "started stream",
    "status": "success"
}
```

##### Stream Upload
Request
```
curl -X POST http://localhost:8080/api/streamupload/vid_7aecfd98-071a-4c62-815c-83e65d95af1d' \
--header 'Content-Type: video/webm' \
--data '@example.webm'
```
Response
```json
{
    "status": "success",
    "message": "Video stream received",
}
```

##### Stop Stream
Request
```
curl -X POST http://localhost:8080/api/endstream/vid_7aecfd98-071a-4c62-815c-83e65d95af1d \
--header 'Content-Type: video/webm' \
--data '@example.webm'
```
Response
```json
{
    "message": "success",
    "data": {
		"video_id":  "vid_7aecfd98-071a-4c62-815c-83e65d95af1d",
		"video_url": "localhost:8080/api/video/vid_7aecfd98-071a-4c62-815c-83e65d95af1d",
	}
}
```

##### View Video Url
Request 
```
curl http://localhost:8080/api/video/vid_7aecfd98-071a-4c62-815c-83e65d95af1d
```
Response: The video file storage location on the disk is rendered

##### View Transcription Url
Request
```
curl http://localhost:8080/api/transcript/vid_7aecfd98-071a-4c62-815c-83e65d95af1d
```
Response: The transcription file storage location for the video is rendered