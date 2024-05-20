package constants

const (
	VideoCreateJSONDecodeErr        = "error in parsing json content for video creation"
	VideoCreateRequestValidationErr = "validation failed for video creation request"
	VideoCreateErr                  = "error occurred in video creation"
	VideoGetIDParamErr              = "error in fetching videoID URL parameter"
	VideoGetFetchErr                = "error occurred in fetching video"
	VideoDeleteIDParamErr           = "error in fetching videoID URL parameter"
	VideoDeleteErr                  = "error occurred in deleting video"

	AnnotationCreateURLParamErr          = "error in fetching videoID URL parameter"
	AnnotationCreateJSONDecodeErr        = "error in parsing json content for annotation creation"
	AnnotationCreateErr                  = "error occurred in video creation"
	AnnotationUpdateVideoIDParamErr      = "error in fetching videoID URL parameter"
	AnnotationUpdateAnnotationIDParamErr = "error in fetching annotationID URL parameter"
	AnnotationUpdateDecodeErr            = "error in parsing json content for annotation update"
	AnnotationUpdateErr                  = "error occurred in updating annotation"
	AnnotationDeleteVideoIDParamErr      = "error in fetching videoID URL parameter"
	AnnotationDeleteAnnotationIDParamErr = "error in fetching annotationID URL parameter"
	AnnotationDeleteErr                  = "error occurred in deleting annotation"

	AnnotationResourceNotFound = "error in fetching annotation resource"
	VideoResourceNotFound      = "error in fetching video resource"

	AnnotationStartTimePositiveErr         = "annotation start time should be positive"
	AnnotationEndTimeGreaterToStartTimeErr = "annotation end time should be greater than start time"
	AnnotationDurationExceedsVideoErr      = "annotation duration exceeds video duration"
	AnnotationTypeEmptyErr                 = "annotation type cannot be empty"
	AnnotationExistsWithSameDurationErr    = "another annotation with same type exists with the same duration"
	AnnotationDoesNotExistForVideoErr      = "annotation does not exists for the video"
	VideoWithSameUrlExistsErr              = "video with Url %s already exists"
)
