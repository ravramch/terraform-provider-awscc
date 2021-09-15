// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The resource with the name requested already exists.
type AlreadyExistsException struct {
	Message *string
}

func (e *AlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AlreadyExistsException) ErrorCode() string             { return "AlreadyExistsException" }
func (e *AlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An error occurred during a CloudFormation registry operation.
type CFNRegistryException struct {
	Message *string
}

func (e *CFNRegistryException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CFNRegistryException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CFNRegistryException) ErrorCode() string             { return "CFNRegistryException" }
func (e *CFNRegistryException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified change set name or ID doesn't exit. To view valid change sets for
// a stack, use the ListChangeSets action.
type ChangeSetNotFoundException struct {
	Message *string
}

func (e *ChangeSetNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ChangeSetNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ChangeSetNotFoundException) ErrorCode() string             { return "ChangeSetNotFoundException" }
func (e *ChangeSetNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ClientTokenConflictException struct {
	Message *string
}

func (e *ClientTokenConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClientTokenConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClientTokenConflictException) ErrorCode() string             { return "ClientTokenConflictException" }
func (e *ClientTokenConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ConcurrentOperationException struct {
	Message *string
}

func (e *ConcurrentOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentOperationException) ErrorCode() string             { return "ConcurrentOperationException" }
func (e *ConcurrentOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified resource exists, but has been changed.
type CreatedButModifiedException struct {
	Message *string
}

func (e *CreatedButModifiedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CreatedButModifiedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CreatedButModifiedException) ErrorCode() string             { return "CreatedButModifiedException" }
func (e *CreatedButModifiedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type GeneralServiceException struct {
	Message *string
}

func (e *GeneralServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *GeneralServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *GeneralServiceException) ErrorCode() string             { return "GeneralServiceException" }
func (e *GeneralServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type HandlerFailureException struct {
	Message *string
}

func (e *HandlerFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *HandlerFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *HandlerFailureException) ErrorCode() string             { return "HandlerFailureException" }
func (e *HandlerFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

type HandlerInternalFailureException struct {
	Message *string
}

func (e *HandlerInternalFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *HandlerInternalFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *HandlerInternalFailureException) ErrorCode() string {
	return "HandlerInternalFailureException"
}
func (e *HandlerInternalFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The template contains resources with capabilities that weren't specified in the
// Capabilities parameter.
type InsufficientCapabilitiesException struct {
	Message *string
}

func (e *InsufficientCapabilitiesException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientCapabilitiesException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientCapabilitiesException) ErrorCode() string {
	return "InsufficientCapabilitiesException"
}
func (e *InsufficientCapabilitiesException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified change set can't be used to update the stack. For example, the
// change set status might be CREATE_IN_PROGRESS, or the stack status might be
// UPDATE_IN_PROGRESS.
type InvalidChangeSetStatusException struct {
	Message *string
}

func (e *InvalidChangeSetStatusException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidChangeSetStatusException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidChangeSetStatusException) ErrorCode() string {
	return "InvalidChangeSetStatusException"
}
func (e *InvalidChangeSetStatusException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type InvalidCredentialsException struct {
	Message *string
}

func (e *InvalidCredentialsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidCredentialsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidCredentialsException) ErrorCode() string             { return "InvalidCredentialsException" }
func (e *InvalidCredentialsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified operation isn't valid.
type InvalidOperationException struct {
	Message *string
}

func (e *InvalidOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidOperationException) ErrorCode() string             { return "InvalidOperationException" }
func (e *InvalidOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type InvalidRequestException struct {
	Message *string
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string             { return "InvalidRequestException" }
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Error reserved for use by the CloudFormation CLI
// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).
// CloudFormation does not return this error to users.
type InvalidStateTransitionException struct {
	Message *string
}

func (e *InvalidStateTransitionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidStateTransitionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidStateTransitionException) ErrorCode() string {
	return "InvalidStateTransitionException"
}
func (e *InvalidStateTransitionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The quota for the resource has already been reached. For information on resource
// and stack limitations, see Limits
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cloudformation-limits.html)
// in the AWS CloudFormation User Guide.
type LimitExceededException struct {
	Message *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified name is already in use.
type NameAlreadyExistsException struct {
	Message *string
}

func (e *NameAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NameAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NameAlreadyExistsException) ErrorCode() string             { return "NameAlreadyExistsException" }
func (e *NameAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type NetworkFailureException struct {
	Message *string
}

func (e *NetworkFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NetworkFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NetworkFailureException) ErrorCode() string             { return "NetworkFailureException" }
func (e *NetworkFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

type NotStabilizedException struct {
	Message *string
}

func (e *NotStabilizedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotStabilizedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotStabilizedException) ErrorCode() string             { return "NotStabilizedException" }
func (e *NotStabilizedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type NotUpdatableException struct {
	Message *string
}

func (e *NotUpdatableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotUpdatableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotUpdatableException) ErrorCode() string             { return "NotUpdatableException" }
func (e *NotUpdatableException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified operation ID already exists.
type OperationIdAlreadyExistsException struct {
	Message *string
}

func (e *OperationIdAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OperationIdAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OperationIdAlreadyExistsException) ErrorCode() string {
	return "OperationIdAlreadyExistsException"
}
func (e *OperationIdAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Another operation is currently in progress for this stack set. Only one
// operation can be performed for a stack set at a given time.
type OperationInProgressException struct {
	Message *string
}

func (e *OperationInProgressException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OperationInProgressException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OperationInProgressException) ErrorCode() string             { return "OperationInProgressException" }
func (e *OperationInProgressException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified ID refers to an operation that doesn't exist.
type OperationNotFoundException struct {
	Message *string
}

func (e *OperationNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OperationNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OperationNotFoundException) ErrorCode() string             { return "OperationNotFoundException" }
func (e *OperationNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Error reserved for use by the CloudFormation CLI
// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).
// CloudFormation does not return this error to users.
type OperationStatusCheckFailedException struct {
	Message *string
}

func (e *OperationStatusCheckFailedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OperationStatusCheckFailedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OperationStatusCheckFailedException) ErrorCode() string {
	return "OperationStatusCheckFailedException"
}
func (e *OperationStatusCheckFailedException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

type PrivateTypeException struct {
	Message *string
}

func (e *PrivateTypeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PrivateTypeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PrivateTypeException) ErrorCode() string             { return "PrivateTypeException" }
func (e *PrivateTypeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type RequestTokenNotFoundException struct {
	Message *string
}

func (e *RequestTokenNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RequestTokenNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RequestTokenNotFoundException) ErrorCode() string             { return "RequestTokenNotFoundException" }
func (e *RequestTokenNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ResourceConflictException struct {
	Message *string
}

func (e *ResourceConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceConflictException) ErrorCode() string             { return "ResourceConflictException" }
func (e *ResourceConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ResourceNotFoundException struct {
	Message *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type ServiceInternalErrorException struct {
	Message *string
}

func (e *ServiceInternalErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceInternalErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceInternalErrorException) ErrorCode() string             { return "ServiceInternalErrorException" }
func (e *ServiceInternalErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

type ServiceLimitExceededException struct {
	Message *string
}

func (e *ServiceLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceLimitExceededException) ErrorCode() string             { return "ServiceLimitExceededException" }
func (e *ServiceLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified stack instance doesn't exist.
type StackInstanceNotFoundException struct {
	Message *string
}

func (e *StackInstanceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StackInstanceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StackInstanceNotFoundException) ErrorCode() string             { return "StackInstanceNotFoundException" }
func (e *StackInstanceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You can't yet delete this stack set, because it still contains one or more stack
// instances. Delete all stack instances from the stack set before deleting the
// stack set.
type StackSetNotEmptyException struct {
	Message *string
}

func (e *StackSetNotEmptyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StackSetNotEmptyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StackSetNotEmptyException) ErrorCode() string             { return "StackSetNotEmptyException" }
func (e *StackSetNotEmptyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified stack set doesn't exist.
type StackSetNotFoundException struct {
	Message *string
}

func (e *StackSetNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StackSetNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StackSetNotFoundException) ErrorCode() string             { return "StackSetNotFoundException" }
func (e *StackSetNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Another operation has been performed on this stack set since the specified
// operation was performed.
type StaleRequestException struct {
	Message *string
}

func (e *StaleRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StaleRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StaleRequestException) ErrorCode() string             { return "StaleRequestException" }
func (e *StaleRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A client request token already exists.
type TokenAlreadyExistsException struct {
	Message *string
}

func (e *TokenAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TokenAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TokenAlreadyExistsException) ErrorCode() string             { return "TokenAlreadyExistsException" }
func (e *TokenAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type TypeConfigurationNotFoundException struct {
	Message *string
}

func (e *TypeConfigurationNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TypeConfigurationNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TypeConfigurationNotFoundException) ErrorCode() string {
	return "TypeConfigurationNotFoundException"
}
func (e *TypeConfigurationNotFoundException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified type does not exist in the CloudFormation registry.
type TypeNotFoundException struct {
	Message *string
}

func (e *TypeNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TypeNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TypeNotFoundException) ErrorCode() string             { return "TypeNotFoundException" }
func (e *TypeNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type UnsupportedActionException struct {
	Message *string
}

func (e *UnsupportedActionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedActionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedActionException) ErrorCode() string             { return "UnsupportedActionException" }
func (e *UnsupportedActionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }