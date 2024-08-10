package constants

// User Errors
const (
	ErrUnauthorized        = "unauthorized"
	ErrInvalidUserID       = "invalid user ID format"
	ErrUserNotFound        = "user not found"
)

// File Upload Errors
const (
	ErrFailedToUploadImage = "failed to upload image"
	ErrFailedToSaveImage   = "failed to save image"
)

// Generic Errors
const (
	ErrInvalidInput = "invalid input data"
)

// Course Errors
const (
	ErrInvalidCourseID       = "invalid course ID"
	ErrCourseNotFound        = "course not found"
	ErrFailedToCreateCourse  = "failed to create course"
	ErrFailedToUpdateCourse  = "failed to update course"
	ErrFailedToDeleteCourse  = "failed to delete course"
	ErrFailedToRetrieveCourse = "failed to retrieve course"
)

// Syllabus Errors
const (
	ErrInvalidSyllabusID       = "invalid syllabus ID"
	ErrSyllabusNotFound        = "syllabus not found"
	ErrFailedToCreateSyllabus  = "failed to create syllabus"
	ErrFailedToUpdateSyllabus  = "failed to update syllabus"
	ErrFailedToDeleteSyllabus  = "failed to delete syllabus"
	ErrUnauthorizedSyllabus    = "unauthorized to modify this syllabus"
	ErrFailedToRetrieveSyllabus = "failed to retrieve syllabus"
)

// Enrollment Errors
const (
	ErrUserAlreadyEnrolled     = "user is already enrolled in the course"
	ErrInsufficientBalance     = "insufficient balance to enroll in the course"
	ErrFailedToCreateEnrollment = "failed to create enrollment"
)

// Forum Post Errors
const (
	ErrFailedToCreateForumPost = "failed to create forum post"
	ErrFailedToRetrievePosts   = "failed to retrieve forum posts"
)

// User Progress Errors
const (
	ErrFailedToUpdateProgress      = "failed to update user progress"
	ErrIncompletePreviousSyllabus  = "complete all previous syllabuses to open this one"
	ErrFailedToRetrieveUserProgress = "failed to retrieve user course progress"
	ErrNoSyllabusesFound           = "no syllabuses found for this course"
)

// Review Errors
const (
	ErrFailedToCreateReview        = "failed to create review"
	ErrFailedToRetrieveReviews     = "failed to retrieve course reviews"
	ErrUserNotEnrolledInCourse     = "user is not enrolled in the course"
	ErrInvalidRate                 = "rate must be between 1 and 5"
	ErrUserAlreadySubmittedReview  = "user has already submitted a review for this course"
)

// Syllabus Material Errors
const (
	ErrInvalidSyllabusMaterialID  = "invalid syllabus material ID"
	ErrUnauthorizedSyllabusAction = "unauthorized to perform this action on the syllabus material"
	ErrFailedToCreateSyllabusMaterial = "failed to create syllabus material"
	ErrFailedToUpdateSyllabusMaterial = "failed to update syllabus material"
	ErrFailedToDeleteSyllabusMaterial = "failed to delete syllabus material"
	ErrFailedToRetrieveSyllabusMaterial = "failed to retrieve syllabus material"
)

// Topup Errors
const (
	ErrFailedToCreateTopup       = "failed to create top-up"
	ErrFailedToUpdateUserBalance = "failed to update user balance"
)


// Assignment Errors
const (
	ErrInvalidAssignmentID        = "invalid assignment ID"
	ErrAssignmentNotFound         = "assignment not found"
	ErrUnauthorizedAssignmentAction = "unauthorized to perform this action on the assignment"
	ErrFailedToCreateAssignment   = "failed to create assignment"
	ErrFailedToUpdateAssignment   = "failed to update assignment"
	ErrFailedToDeleteAssignment   = "failed to delete assignment"
	ErrFailedToRetrieveAssignment = "failed to retrieve assignment"
	ErrFailedToCreateUserAssignment = "failed to create user assignment"
)