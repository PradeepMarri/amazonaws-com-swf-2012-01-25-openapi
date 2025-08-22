package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Decision represents the Decision schema from the OpenAPI specification
type Decision struct {
	Completeworkflowexecutiondecisionattributes interface{} `json:"completeWorkflowExecutionDecisionAttributes,omitempty"`
	Recordmarkerdecisionattributes interface{} `json:"recordMarkerDecisionAttributes,omitempty"`
	Requestcancelactivitytaskdecisionattributes interface{} `json:"requestCancelActivityTaskDecisionAttributes,omitempty"`
	Decisiontype interface{} `json:"decisionType"`
	Failworkflowexecutiondecisionattributes interface{} `json:"failWorkflowExecutionDecisionAttributes,omitempty"`
	Requestcancelexternalworkflowexecutiondecisionattributes interface{} `json:"requestCancelExternalWorkflowExecutionDecisionAttributes,omitempty"`
	Scheduleactivitytaskdecisionattributes interface{} `json:"scheduleActivityTaskDecisionAttributes,omitempty"`
	Signalexternalworkflowexecutiondecisionattributes interface{} `json:"signalExternalWorkflowExecutionDecisionAttributes,omitempty"`
	Startchildworkflowexecutiondecisionattributes interface{} `json:"startChildWorkflowExecutionDecisionAttributes,omitempty"`
	Starttimerdecisionattributes interface{} `json:"startTimerDecisionAttributes,omitempty"`
	Canceltimerdecisionattributes interface{} `json:"cancelTimerDecisionAttributes,omitempty"`
	Continueasnewworkflowexecutiondecisionattributes interface{} `json:"continueAsNewWorkflowExecutionDecisionAttributes,omitempty"`
	Schedulelambdafunctiondecisionattributes interface{} `json:"scheduleLambdaFunctionDecisionAttributes,omitempty"`
	Cancelworkflowexecutiondecisionattributes interface{} `json:"cancelWorkflowExecutionDecisionAttributes,omitempty"`
}

// ActivityTypeInfos represents the ActivityTypeInfos schema from the OpenAPI specification
type ActivityTypeInfos struct {
	Nextpagetoken interface{} `json:"nextPageToken,omitempty"`
	Typeinfos interface{} `json:"typeInfos"`
}

// ListTagsForResourceInput represents the ListTagsForResourceInput schema from the OpenAPI specification
type ListTagsForResourceInput struct {
	Resourcearn interface{} `json:"resourceArn"`
}

// ContinueAsNewWorkflowExecutionFailedEventAttributes represents the ContinueAsNewWorkflowExecutionFailedEventAttributes schema from the OpenAPI specification
type ContinueAsNewWorkflowExecutionFailedEventAttributes struct {
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Cause interface{} `json:"cause"`
}

// History represents the History schema from the OpenAPI specification
type History struct {
	Events interface{} `json:"events"`
	Nextpagetoken interface{} `json:"nextPageToken,omitempty"`
}

// PendingTaskCount represents the PendingTaskCount schema from the OpenAPI specification
type PendingTaskCount struct {
	Count interface{} `json:"count"`
	Truncated interface{} `json:"truncated,omitempty"`
}

// UndeprecateActivityTypeInput represents the UndeprecateActivityTypeInput schema from the OpenAPI specification
type UndeprecateActivityTypeInput struct {
	Activitytype interface{} `json:"activityType"`
	Domain interface{} `json:"domain"`
}

// ActivityTaskTimedOutEventAttributes represents the ActivityTaskTimedOutEventAttributes schema from the OpenAPI specification
type ActivityTaskTimedOutEventAttributes struct {
	Details interface{} `json:"details,omitempty"`
	Scheduledeventid interface{} `json:"scheduledEventId"`
	Startedeventid interface{} `json:"startedEventId"`
	Timeouttype interface{} `json:"timeoutType"`
}

// WorkflowExecutionConfiguration represents the WorkflowExecutionConfiguration schema from the OpenAPI specification
type WorkflowExecutionConfiguration struct {
	Taskpriority interface{} `json:"taskPriority,omitempty"`
	Taskstarttoclosetimeout interface{} `json:"taskStartToCloseTimeout"`
	Childpolicy interface{} `json:"childPolicy"`
	Executionstarttoclosetimeout interface{} `json:"executionStartToCloseTimeout"`
	Lambdarole interface{} `json:"lambdaRole,omitempty"`
	Tasklist interface{} `json:"taskList"`
}

// RecordMarkerDecisionAttributes represents the RecordMarkerDecisionAttributes schema from the OpenAPI specification
type RecordMarkerDecisionAttributes struct {
	Details interface{} `json:"details,omitempty"`
	Markername interface{} `json:"markerName"`
}

// RequestCancelActivityTaskDecisionAttributes represents the RequestCancelActivityTaskDecisionAttributes schema from the OpenAPI specification
type RequestCancelActivityTaskDecisionAttributes struct {
	Activityid interface{} `json:"activityId"`
}

// ExternalWorkflowExecutionSignaledEventAttributes represents the ExternalWorkflowExecutionSignaledEventAttributes schema from the OpenAPI specification
type ExternalWorkflowExecutionSignaledEventAttributes struct {
	Initiatedeventid interface{} `json:"initiatedEventId"`
	Workflowexecution interface{} `json:"workflowExecution"`
}

// TimerCanceledEventAttributes represents the TimerCanceledEventAttributes schema from the OpenAPI specification
type TimerCanceledEventAttributes struct {
	Timerid interface{} `json:"timerId"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Startedeventid interface{} `json:"startedEventId"`
}

// DescribeActivityTypeInput represents the DescribeActivityTypeInput schema from the OpenAPI specification
type DescribeActivityTypeInput struct {
	Activitytype interface{} `json:"activityType"`
	Domain interface{} `json:"domain"`
}

// DomainDetail represents the DomainDetail schema from the OpenAPI specification
type DomainDetail struct {
	Domaininfo interface{} `json:"domainInfo"`
	Configuration interface{} `json:"configuration"`
}

// RequestCancelExternalWorkflowExecutionFailedEventAttributes represents the RequestCancelExternalWorkflowExecutionFailedEventAttributes schema from the OpenAPI specification
type RequestCancelExternalWorkflowExecutionFailedEventAttributes struct {
	Cause interface{} `json:"cause"`
	Control interface{} `json:"control,omitempty"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Initiatedeventid interface{} `json:"initiatedEventId"`
	Runid interface{} `json:"runId,omitempty"`
	Workflowid interface{} `json:"workflowId"`
}

// UndeprecateWorkflowTypeInput represents the UndeprecateWorkflowTypeInput schema from the OpenAPI specification
type UndeprecateWorkflowTypeInput struct {
	Workflowtype interface{} `json:"workflowType"`
	Domain interface{} `json:"domain"`
}

// PollForActivityTaskInput represents the PollForActivityTaskInput schema from the OpenAPI specification
type PollForActivityTaskInput struct {
	Identity interface{} `json:"identity,omitempty"`
	Tasklist interface{} `json:"taskList"`
	Domain interface{} `json:"domain"`
}

// RequestCancelActivityTaskFailedEventAttributes represents the RequestCancelActivityTaskFailedEventAttributes schema from the OpenAPI specification
type RequestCancelActivityTaskFailedEventAttributes struct {
	Cause interface{} `json:"cause"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Activityid interface{} `json:"activityId"`
}

// WorkflowTypeFilter represents the WorkflowTypeFilter schema from the OpenAPI specification
type WorkflowTypeFilter struct {
	Name interface{} `json:"name"`
	Version interface{} `json:"version,omitempty"`
}

// LambdaFunctionFailedEventAttributes represents the LambdaFunctionFailedEventAttributes schema from the OpenAPI specification
type LambdaFunctionFailedEventAttributes struct {
	Reason interface{} `json:"reason,omitempty"`
	Scheduledeventid interface{} `json:"scheduledEventId"`
	Startedeventid interface{} `json:"startedEventId"`
	Details interface{} `json:"details,omitempty"`
}

// WorkflowExecutionInfos represents the WorkflowExecutionInfos schema from the OpenAPI specification
type WorkflowExecutionInfos struct {
	Nextpagetoken interface{} `json:"nextPageToken,omitempty"`
	Executioninfos interface{} `json:"executionInfos"`
}

// DomainInfo represents the DomainInfo schema from the OpenAPI specification
type DomainInfo struct {
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name"`
	Status interface{} `json:"status"`
	Arn interface{} `json:"arn,omitempty"`
}

// StartLambdaFunctionFailedEventAttributes represents the StartLambdaFunctionFailedEventAttributes schema from the OpenAPI specification
type StartLambdaFunctionFailedEventAttributes struct {
	Cause interface{} `json:"cause,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Scheduledeventid interface{} `json:"scheduledEventId,omitempty"`
}

// RegisterWorkflowTypeInput represents the RegisterWorkflowTypeInput schema from the OpenAPI specification
type RegisterWorkflowTypeInput struct {
	Name interface{} `json:"name"`
	Defaultlambdarole interface{} `json:"defaultLambdaRole,omitempty"`
	Domain interface{} `json:"domain"`
	Defaultchildpolicy interface{} `json:"defaultChildPolicy,omitempty"`
	Defaultexecutionstarttoclosetimeout interface{} `json:"defaultExecutionStartToCloseTimeout,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Version interface{} `json:"version"`
	Defaulttasklist interface{} `json:"defaultTaskList,omitempty"`
	Defaulttaskpriority interface{} `json:"defaultTaskPriority,omitempty"`
	Defaulttaskstarttoclosetimeout interface{} `json:"defaultTaskStartToCloseTimeout,omitempty"`
}

// PollForDecisionTaskInput represents the PollForDecisionTaskInput schema from the OpenAPI specification
type PollForDecisionTaskInput struct {
	Nextpagetoken interface{} `json:"nextPageToken,omitempty"`
	Reverseorder interface{} `json:"reverseOrder,omitempty"`
	Startatpreviousstartedevent interface{} `json:"startAtPreviousStartedEvent,omitempty"`
	Tasklist interface{} `json:"taskList"`
	Domain interface{} `json:"domain"`
	Identity interface{} `json:"identity,omitempty"`
	Maximumpagesize interface{} `json:"maximumPageSize,omitempty"`
}

// UndeprecateDomainInput represents the UndeprecateDomainInput schema from the OpenAPI specification
type UndeprecateDomainInput struct {
	Name interface{} `json:"name"`
}

// SignalExternalWorkflowExecutionFailedEventAttributes represents the SignalExternalWorkflowExecutionFailedEventAttributes schema from the OpenAPI specification
type SignalExternalWorkflowExecutionFailedEventAttributes struct {
	Control interface{} `json:"control,omitempty"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Initiatedeventid interface{} `json:"initiatedEventId"`
	Runid interface{} `json:"runId,omitempty"`
	Workflowid interface{} `json:"workflowId"`
	Cause interface{} `json:"cause"`
}

// StartChildWorkflowExecutionInitiatedEventAttributes represents the StartChildWorkflowExecutionInitiatedEventAttributes schema from the OpenAPI specification
type StartChildWorkflowExecutionInitiatedEventAttributes struct {
	Workflowid interface{} `json:"workflowId"`
	Childpolicy interface{} `json:"childPolicy"`
	Lambdarole interface{} `json:"lambdaRole,omitempty"`
	Taskstarttoclosetimeout interface{} `json:"taskStartToCloseTimeout,omitempty"`
	Taglist interface{} `json:"tagList,omitempty"`
	Workflowtype interface{} `json:"workflowType"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Input interface{} `json:"input,omitempty"`
	Tasklist interface{} `json:"taskList"`
	Control interface{} `json:"control,omitempty"`
	Executionstarttoclosetimeout interface{} `json:"executionStartToCloseTimeout,omitempty"`
	Taskpriority interface{} `json:"taskPriority,omitempty"`
}

// CancelTimerFailedEventAttributes represents the CancelTimerFailedEventAttributes schema from the OpenAPI specification
type CancelTimerFailedEventAttributes struct {
	Cause interface{} `json:"cause"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Timerid interface{} `json:"timerId"`
}

// CountClosedWorkflowExecutionsInput represents the CountClosedWorkflowExecutionsInput schema from the OpenAPI specification
type CountClosedWorkflowExecutionsInput struct {
	Tagfilter interface{} `json:"tagFilter,omitempty"`
	Typefilter interface{} `json:"typeFilter,omitempty"`
	Closestatusfilter interface{} `json:"closeStatusFilter,omitempty"`
	Closetimefilter interface{} `json:"closeTimeFilter,omitempty"`
	Domain interface{} `json:"domain"`
	Executionfilter interface{} `json:"executionFilter,omitempty"`
	Starttimefilter interface{} `json:"startTimeFilter,omitempty"`
}

// ListWorkflowTypesInput represents the ListWorkflowTypesInput schema from the OpenAPI specification
type ListWorkflowTypesInput struct {
	Nextpagetoken interface{} `json:"nextPageToken,omitempty"`
	Registrationstatus interface{} `json:"registrationStatus"`
	Reverseorder interface{} `json:"reverseOrder,omitempty"`
	Domain interface{} `json:"domain"`
	Maximumpagesize interface{} `json:"maximumPageSize,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// DecisionTask represents the DecisionTask schema from the OpenAPI specification
type DecisionTask struct {
	Workflowexecution interface{} `json:"workflowExecution"`
	Workflowtype interface{} `json:"workflowType"`
	Events interface{} `json:"events"`
	Nextpagetoken interface{} `json:"nextPageToken,omitempty"`
	Previousstartedeventid interface{} `json:"previousStartedEventId,omitempty"`
	Startedeventid interface{} `json:"startedEventId"`
	Tasktoken interface{} `json:"taskToken"`
}

// ChildWorkflowExecutionCanceledEventAttributes represents the ChildWorkflowExecutionCanceledEventAttributes schema from the OpenAPI specification
type ChildWorkflowExecutionCanceledEventAttributes struct {
	Workflowtype interface{} `json:"workflowType"`
	Details interface{} `json:"details,omitempty"`
	Initiatedeventid interface{} `json:"initiatedEventId"`
	Startedeventid interface{} `json:"startedEventId"`
	Workflowexecution interface{} `json:"workflowExecution"`
}

// TagResourceInput represents the TagResourceInput schema from the OpenAPI specification
type TagResourceInput struct {
	Resourcearn interface{} `json:"resourceArn"`
	Tags interface{} `json:"tags"`
}

// TimerStartedEventAttributes represents the TimerStartedEventAttributes schema from the OpenAPI specification
type TimerStartedEventAttributes struct {
	Timerid interface{} `json:"timerId"`
	Control interface{} `json:"control,omitempty"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Starttofiretimeout interface{} `json:"startToFireTimeout"`
}

// WorkflowExecutionDetail represents the WorkflowExecutionDetail schema from the OpenAPI specification
type WorkflowExecutionDetail struct {
	Opencounts interface{} `json:"openCounts"`
	Executionconfiguration interface{} `json:"executionConfiguration"`
	Executioninfo interface{} `json:"executionInfo"`
	Latestactivitytasktimestamp interface{} `json:"latestActivityTaskTimestamp,omitempty"`
	Latestexecutioncontext interface{} `json:"latestExecutionContext,omitempty"`
}

// DescribeWorkflowExecutionInput represents the DescribeWorkflowExecutionInput schema from the OpenAPI specification
type DescribeWorkflowExecutionInput struct {
	Domain interface{} `json:"domain"`
	Execution interface{} `json:"execution"`
}

// ListTagsForResourceOutput represents the ListTagsForResourceOutput schema from the OpenAPI specification
type ListTagsForResourceOutput struct {
	Tags interface{} `json:"tags,omitempty"`
}

// SignalExternalWorkflowExecutionInitiatedEventAttributes represents the SignalExternalWorkflowExecutionInitiatedEventAttributes schema from the OpenAPI specification
type SignalExternalWorkflowExecutionInitiatedEventAttributes struct {
	Control interface{} `json:"control,omitempty"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Input interface{} `json:"input,omitempty"`
	Runid interface{} `json:"runId,omitempty"`
	Signalname interface{} `json:"signalName"`
	Workflowid interface{} `json:"workflowId"`
}

// ScheduleLambdaFunctionFailedEventAttributes represents the ScheduleLambdaFunctionFailedEventAttributes schema from the OpenAPI specification
type ScheduleLambdaFunctionFailedEventAttributes struct {
	Cause interface{} `json:"cause"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Id interface{} `json:"id"`
	Name interface{} `json:"name"`
}

// CountPendingActivityTasksInput represents the CountPendingActivityTasksInput schema from the OpenAPI specification
type CountPendingActivityTasksInput struct {
	Tasklist interface{} `json:"taskList"`
	Domain interface{} `json:"domain"`
}

// DescribeDomainInput represents the DescribeDomainInput schema from the OpenAPI specification
type DescribeDomainInput struct {
	Name interface{} `json:"name"`
}

// ActivityTaskFailedEventAttributes represents the ActivityTaskFailedEventAttributes schema from the OpenAPI specification
type ActivityTaskFailedEventAttributes struct {
	Details interface{} `json:"details,omitempty"`
	Reason interface{} `json:"reason,omitempty"`
	Scheduledeventid interface{} `json:"scheduledEventId"`
	Startedeventid interface{} `json:"startedEventId"`
}

// CountPendingDecisionTasksInput represents the CountPendingDecisionTasksInput schema from the OpenAPI specification
type CountPendingDecisionTasksInput struct {
	Domain interface{} `json:"domain"`
	Tasklist interface{} `json:"taskList"`
}

// RequestCancelExternalWorkflowExecutionInitiatedEventAttributes represents the RequestCancelExternalWorkflowExecutionInitiatedEventAttributes schema from the OpenAPI specification
type RequestCancelExternalWorkflowExecutionInitiatedEventAttributes struct {
	Control interface{} `json:"control,omitempty"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Runid interface{} `json:"runId,omitempty"`
	Workflowid interface{} `json:"workflowId"`
}

// WorkflowTypeInfo represents the WorkflowTypeInfo schema from the OpenAPI specification
type WorkflowTypeInfo struct {
	Creationdate interface{} `json:"creationDate"`
	Deprecationdate interface{} `json:"deprecationDate,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Status interface{} `json:"status"`
	Workflowtype interface{} `json:"workflowType"`
}

// ActivityTask represents the ActivityTask schema from the OpenAPI specification
type ActivityTask struct {
	Tasktoken interface{} `json:"taskToken"`
	Workflowexecution interface{} `json:"workflowExecution"`
	Activityid interface{} `json:"activityId"`
	Activitytype interface{} `json:"activityType"`
	Input interface{} `json:"input,omitempty"`
	Startedeventid interface{} `json:"startedEventId"`
}

// DeprecateWorkflowTypeInput represents the DeprecateWorkflowTypeInput schema from the OpenAPI specification
type DeprecateWorkflowTypeInput struct {
	Domain interface{} `json:"domain"`
	Workflowtype interface{} `json:"workflowType"`
}

// LambdaFunctionTimedOutEventAttributes represents the LambdaFunctionTimedOutEventAttributes schema from the OpenAPI specification
type LambdaFunctionTimedOutEventAttributes struct {
	Startedeventid interface{} `json:"startedEventId"`
	Timeouttype interface{} `json:"timeoutType,omitempty"`
	Scheduledeventid interface{} `json:"scheduledEventId"`
}

// ActivityTypeConfiguration represents the ActivityTypeConfiguration schema from the OpenAPI specification
type ActivityTypeConfiguration struct {
	Defaulttaskpriority interface{} `json:"defaultTaskPriority,omitempty"`
	Defaulttaskscheduletoclosetimeout interface{} `json:"defaultTaskScheduleToCloseTimeout,omitempty"`
	Defaulttaskscheduletostarttimeout interface{} `json:"defaultTaskScheduleToStartTimeout,omitempty"`
	Defaulttaskstarttoclosetimeout interface{} `json:"defaultTaskStartToCloseTimeout,omitempty"`
	Defaulttaskheartbeattimeout interface{} `json:"defaultTaskHeartbeatTimeout,omitempty"`
	Defaulttasklist interface{} `json:"defaultTaskList,omitempty"`
}

// ScheduleActivityTaskDecisionAttributes represents the ScheduleActivityTaskDecisionAttributes schema from the OpenAPI specification
type ScheduleActivityTaskDecisionAttributes struct {
	Tasklist interface{} `json:"taskList,omitempty"`
	Taskpriority interface{} `json:"taskPriority,omitempty"`
	Scheduletoclosetimeout interface{} `json:"scheduleToCloseTimeout,omitempty"`
	Heartbeattimeout interface{} `json:"heartbeatTimeout,omitempty"`
	Input interface{} `json:"input,omitempty"`
	Scheduletostarttimeout interface{} `json:"scheduleToStartTimeout,omitempty"`
	Activitytype interface{} `json:"activityType"`
	Control interface{} `json:"control,omitempty"`
	Activityid interface{} `json:"activityId"`
	Starttoclosetimeout interface{} `json:"startToCloseTimeout,omitempty"`
}

// DomainInfos represents the DomainInfos schema from the OpenAPI specification
type DomainInfos struct {
	Domaininfos interface{} `json:"domainInfos"`
	Nextpagetoken interface{} `json:"nextPageToken,omitempty"`
}

// CancelWorkflowExecutionFailedEventAttributes represents the CancelWorkflowExecutionFailedEventAttributes schema from the OpenAPI specification
type CancelWorkflowExecutionFailedEventAttributes struct {
	Cause interface{} `json:"cause"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
}

// ExternalWorkflowExecutionCancelRequestedEventAttributes represents the ExternalWorkflowExecutionCancelRequestedEventAttributes schema from the OpenAPI specification
type ExternalWorkflowExecutionCancelRequestedEventAttributes struct {
	Initiatedeventid interface{} `json:"initiatedEventId"`
	Workflowexecution interface{} `json:"workflowExecution"`
}

// WorkflowExecutionCount represents the WorkflowExecutionCount schema from the OpenAPI specification
type WorkflowExecutionCount struct {
	Count interface{} `json:"count"`
	Truncated interface{} `json:"truncated,omitempty"`
}

// RequestCancelExternalWorkflowExecutionDecisionAttributes represents the RequestCancelExternalWorkflowExecutionDecisionAttributes schema from the OpenAPI specification
type RequestCancelExternalWorkflowExecutionDecisionAttributes struct {
	Control interface{} `json:"control,omitempty"`
	Runid interface{} `json:"runId,omitempty"`
	Workflowid interface{} `json:"workflowId"`
}

// ActivityTaskCancelRequestedEventAttributes represents the ActivityTaskCancelRequestedEventAttributes schema from the OpenAPI specification
type ActivityTaskCancelRequestedEventAttributes struct {
	Activityid interface{} `json:"activityId"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
}

// MarkerRecordedEventAttributes represents the MarkerRecordedEventAttributes schema from the OpenAPI specification
type MarkerRecordedEventAttributes struct {
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Details interface{} `json:"details,omitempty"`
	Markername interface{} `json:"markerName"`
}

// StartWorkflowExecutionInput represents the StartWorkflowExecutionInput schema from the OpenAPI specification
type StartWorkflowExecutionInput struct {
	Input interface{} `json:"input,omitempty"`
	Taskstarttoclosetimeout interface{} `json:"taskStartToCloseTimeout,omitempty"`
	Taglist interface{} `json:"tagList,omitempty"`
	Taskpriority interface{} `json:"taskPriority,omitempty"`
	Domain interface{} `json:"domain"`
	Executionstarttoclosetimeout interface{} `json:"executionStartToCloseTimeout,omitempty"`
	Lambdarole interface{} `json:"lambdaRole,omitempty"`
	Tasklist interface{} `json:"taskList,omitempty"`
	Workflowid interface{} `json:"workflowId"`
	Workflowtype interface{} `json:"workflowType"`
	Childpolicy interface{} `json:"childPolicy,omitempty"`
}

// WorkflowExecutionContinuedAsNewEventAttributes represents the WorkflowExecutionContinuedAsNewEventAttributes schema from the OpenAPI specification
type WorkflowExecutionContinuedAsNewEventAttributes struct {
	Tasklist interface{} `json:"taskList"`
	Childpolicy interface{} `json:"childPolicy"`
	Lambdarole interface{} `json:"lambdaRole,omitempty"`
	Taskpriority interface{} `json:"taskPriority,omitempty"`
	Taskstarttoclosetimeout interface{} `json:"taskStartToCloseTimeout,omitempty"`
	Workflowtype interface{} `json:"workflowType"`
	Input interface{} `json:"input,omitempty"`
	Newexecutionrunid interface{} `json:"newExecutionRunId"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Executionstarttoclosetimeout interface{} `json:"executionStartToCloseTimeout,omitempty"`
	Taglist interface{} `json:"tagList,omitempty"`
}

// WorkflowType represents the WorkflowType schema from the OpenAPI specification
type WorkflowType struct {
	Name interface{} `json:"name"`
	Version interface{} `json:"version"`
}

// ActivityType represents the ActivityType schema from the OpenAPI specification
type ActivityType struct {
	Name interface{} `json:"name"`
	Version interface{} `json:"version"`
}

// GetWorkflowExecutionHistoryInput represents the GetWorkflowExecutionHistoryInput schema from the OpenAPI specification
type GetWorkflowExecutionHistoryInput struct {
	Domain interface{} `json:"domain"`
	Execution interface{} `json:"execution"`
	Maximumpagesize interface{} `json:"maximumPageSize,omitempty"`
	Nextpagetoken interface{} `json:"nextPageToken,omitempty"`
	Reverseorder interface{} `json:"reverseOrder,omitempty"`
}

// RecordActivityTaskHeartbeatInput represents the RecordActivityTaskHeartbeatInput schema from the OpenAPI specification
type RecordActivityTaskHeartbeatInput struct {
	Details interface{} `json:"details,omitempty"`
	Tasktoken interface{} `json:"taskToken"`
}

// WorkflowExecutionStartedEventAttributes represents the WorkflowExecutionStartedEventAttributes schema from the OpenAPI specification
type WorkflowExecutionStartedEventAttributes struct {
	Workflowtype interface{} `json:"workflowType"`
	Childpolicy interface{} `json:"childPolicy"`
	Continuedexecutionrunid interface{} `json:"continuedExecutionRunId,omitempty"`
	Input interface{} `json:"input,omitempty"`
	Lambdarole interface{} `json:"lambdaRole,omitempty"`
	Parentinitiatedeventid interface{} `json:"parentInitiatedEventId,omitempty"`
	Tasklist interface{} `json:"taskList"`
	Taskpriority interface{} `json:"taskPriority,omitempty"`
	Executionstarttoclosetimeout interface{} `json:"executionStartToCloseTimeout,omitempty"`
	Taskstarttoclosetimeout interface{} `json:"taskStartToCloseTimeout,omitempty"`
	Parentworkflowexecution interface{} `json:"parentWorkflowExecution,omitempty"`
	Taglist interface{} `json:"tagList,omitempty"`
}

// Run represents the Run schema from the OpenAPI specification
type Run struct {
	Runid interface{} `json:"runId,omitempty"`
}

// ScheduleLambdaFunctionDecisionAttributes represents the ScheduleLambdaFunctionDecisionAttributes schema from the OpenAPI specification
type ScheduleLambdaFunctionDecisionAttributes struct {
	Name interface{} `json:"name"`
	Starttoclosetimeout interface{} `json:"startToCloseTimeout,omitempty"`
	Control interface{} `json:"control,omitempty"`
	Id interface{} `json:"id"`
	Input interface{} `json:"input,omitempty"`
}

// ListClosedWorkflowExecutionsInput represents the ListClosedWorkflowExecutionsInput schema from the OpenAPI specification
type ListClosedWorkflowExecutionsInput struct {
	Executionfilter interface{} `json:"executionFilter,omitempty"`
	Tagfilter interface{} `json:"tagFilter,omitempty"`
	Nextpagetoken interface{} `json:"nextPageToken,omitempty"`
	Starttimefilter interface{} `json:"startTimeFilter,omitempty"`
	Closestatusfilter interface{} `json:"closeStatusFilter,omitempty"`
	Maximumpagesize interface{} `json:"maximumPageSize,omitempty"`
	Typefilter interface{} `json:"typeFilter,omitempty"`
	Closetimefilter interface{} `json:"closeTimeFilter,omitempty"`
	Reverseorder interface{} `json:"reverseOrder,omitempty"`
	Domain interface{} `json:"domain"`
}

// ChildWorkflowExecutionTerminatedEventAttributes represents the ChildWorkflowExecutionTerminatedEventAttributes schema from the OpenAPI specification
type ChildWorkflowExecutionTerminatedEventAttributes struct {
	Workflowexecution interface{} `json:"workflowExecution"`
	Workflowtype interface{} `json:"workflowType"`
	Initiatedeventid interface{} `json:"initiatedEventId"`
	Startedeventid interface{} `json:"startedEventId"`
}

// ActivityTaskCompletedEventAttributes represents the ActivityTaskCompletedEventAttributes schema from the OpenAPI specification
type ActivityTaskCompletedEventAttributes struct {
	Startedeventid interface{} `json:"startedEventId"`
	Result interface{} `json:"result,omitempty"`
	Scheduledeventid interface{} `json:"scheduledEventId"`
}

// TimerFiredEventAttributes represents the TimerFiredEventAttributes schema from the OpenAPI specification
type TimerFiredEventAttributes struct {
	Startedeventid interface{} `json:"startedEventId"`
	Timerid interface{} `json:"timerId"`
}

// CancelTimerDecisionAttributes represents the CancelTimerDecisionAttributes schema from the OpenAPI specification
type CancelTimerDecisionAttributes struct {
	Timerid interface{} `json:"timerId"`
}

// ChildWorkflowExecutionStartedEventAttributes represents the ChildWorkflowExecutionStartedEventAttributes schema from the OpenAPI specification
type ChildWorkflowExecutionStartedEventAttributes struct {
	Initiatedeventid interface{} `json:"initiatedEventId"`
	Workflowexecution interface{} `json:"workflowExecution"`
	Workflowtype interface{} `json:"workflowType"`
}

// WorkflowExecutionFilter represents the WorkflowExecutionFilter schema from the OpenAPI specification
type WorkflowExecutionFilter struct {
	Workflowid interface{} `json:"workflowId"`
}

// WorkflowTypeConfiguration represents the WorkflowTypeConfiguration schema from the OpenAPI specification
type WorkflowTypeConfiguration struct {
	Defaulttaskstarttoclosetimeout interface{} `json:"defaultTaskStartToCloseTimeout,omitempty"`
	Defaultchildpolicy interface{} `json:"defaultChildPolicy,omitempty"`
	Defaultexecutionstarttoclosetimeout interface{} `json:"defaultExecutionStartToCloseTimeout,omitempty"`
	Defaultlambdarole interface{} `json:"defaultLambdaRole,omitempty"`
	Defaulttasklist interface{} `json:"defaultTaskList,omitempty"`
	Defaulttaskpriority interface{} `json:"defaultTaskPriority,omitempty"`
}

// ActivityTypeInfo represents the ActivityTypeInfo schema from the OpenAPI specification
type ActivityTypeInfo struct {
	Deprecationdate interface{} `json:"deprecationDate,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Status interface{} `json:"status"`
	Activitytype interface{} `json:"activityType"`
	Creationdate interface{} `json:"creationDate"`
}

// CloseStatusFilter represents the CloseStatusFilter schema from the OpenAPI specification
type CloseStatusFilter struct {
	Status interface{} `json:"status"`
}

// ActivityTaskCanceledEventAttributes represents the ActivityTaskCanceledEventAttributes schema from the OpenAPI specification
type ActivityTaskCanceledEventAttributes struct {
	Startedeventid interface{} `json:"startedEventId"`
	Details interface{} `json:"details,omitempty"`
	Latestcancelrequestedeventid interface{} `json:"latestCancelRequestedEventId,omitempty"`
	Scheduledeventid interface{} `json:"scheduledEventId"`
}

// DecisionTaskScheduledEventAttributes represents the DecisionTaskScheduledEventAttributes schema from the OpenAPI specification
type DecisionTaskScheduledEventAttributes struct {
	Tasklist interface{} `json:"taskList"`
	Taskpriority interface{} `json:"taskPriority,omitempty"`
	Starttoclosetimeout interface{} `json:"startToCloseTimeout,omitempty"`
}

// LambdaFunctionStartedEventAttributes represents the LambdaFunctionStartedEventAttributes schema from the OpenAPI specification
type LambdaFunctionStartedEventAttributes struct {
	Scheduledeventid interface{} `json:"scheduledEventId"`
}

// UntagResourceInput represents the UntagResourceInput schema from the OpenAPI specification
type UntagResourceInput struct {
	Resourcearn interface{} `json:"resourceArn"`
	Tagkeys interface{} `json:"tagKeys"`
}

// WorkflowExecution represents the WorkflowExecution schema from the OpenAPI specification
type WorkflowExecution struct {
	Runid interface{} `json:"runId"`
	Workflowid interface{} `json:"workflowId"`
}

// ResourceTag represents the ResourceTag schema from the OpenAPI specification
type ResourceTag struct {
	Key interface{} `json:"key"`
	Value interface{} `json:"value,omitempty"`
}

// WorkflowExecutionOpenCounts represents the WorkflowExecutionOpenCounts schema from the OpenAPI specification
type WorkflowExecutionOpenCounts struct {
	Openlambdafunctions interface{} `json:"openLambdaFunctions,omitempty"`
	Opentimers interface{} `json:"openTimers"`
	Openactivitytasks interface{} `json:"openActivityTasks"`
	Openchildworkflowexecutions interface{} `json:"openChildWorkflowExecutions"`
	Opendecisiontasks interface{} `json:"openDecisionTasks"`
}

// RespondActivityTaskCanceledInput represents the RespondActivityTaskCanceledInput schema from the OpenAPI specification
type RespondActivityTaskCanceledInput struct {
	Tasktoken interface{} `json:"taskToken"`
	Details interface{} `json:"details,omitempty"`
}

// LambdaFunctionScheduledEventAttributes represents the LambdaFunctionScheduledEventAttributes schema from the OpenAPI specification
type LambdaFunctionScheduledEventAttributes struct {
	Control interface{} `json:"control,omitempty"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Id interface{} `json:"id"`
	Input interface{} `json:"input,omitempty"`
	Name interface{} `json:"name"`
	Starttoclosetimeout interface{} `json:"startToCloseTimeout,omitempty"`
}

// WorkflowTypeInfos represents the WorkflowTypeInfos schema from the OpenAPI specification
type WorkflowTypeInfos struct {
	Nextpagetoken interface{} `json:"nextPageToken,omitempty"`
	Typeinfos interface{} `json:"typeInfos"`
}

// WorkflowExecutionTimedOutEventAttributes represents the WorkflowExecutionTimedOutEventAttributes schema from the OpenAPI specification
type WorkflowExecutionTimedOutEventAttributes struct {
	Childpolicy interface{} `json:"childPolicy"`
	Timeouttype interface{} `json:"timeoutType"`
}

// HistoryEvent represents the HistoryEvent schema from the OpenAPI specification
type HistoryEvent struct {
	Eventid interface{} `json:"eventId"`
	Workflowexecutionstartedeventattributes interface{} `json:"workflowExecutionStartedEventAttributes,omitempty"`
	Eventtimestamp interface{} `json:"eventTimestamp"`
	Workflowexecutiontimedouteventattributes interface{} `json:"workflowExecutionTimedOutEventAttributes,omitempty"`
	Decisiontasktimedouteventattributes interface{} `json:"decisionTaskTimedOutEventAttributes,omitempty"`
	Workflowexecutioncontinuedasneweventattributes interface{} `json:"workflowExecutionContinuedAsNewEventAttributes,omitempty"`
	Startchildworkflowexecutioninitiatedeventattributes interface{} `json:"startChildWorkflowExecutionInitiatedEventAttributes,omitempty"`
	Timerfiredeventattributes interface{} `json:"timerFiredEventAttributes,omitempty"`
	Schedulelambdafunctionfailedeventattributes interface{} `json:"scheduleLambdaFunctionFailedEventAttributes,omitempty"`
	Decisiontaskcompletedeventattributes interface{} `json:"decisionTaskCompletedEventAttributes,omitempty"`
	Workflowexecutioncanceledeventattributes interface{} `json:"workflowExecutionCanceledEventAttributes,omitempty"`
	Continueasnewworkflowexecutionfailedeventattributes interface{} `json:"continueAsNewWorkflowExecutionFailedEventAttributes,omitempty"`
	Signalexternalworkflowexecutioninitiatedeventattributes interface{} `json:"signalExternalWorkflowExecutionInitiatedEventAttributes,omitempty"`
	Recordmarkerfailedeventattributes interface{} `json:"recordMarkerFailedEventAttributes,omitempty"`
	Childworkflowexecutionterminatedeventattributes interface{} `json:"childWorkflowExecutionTerminatedEventAttributes,omitempty"`
	Decisiontaskstartedeventattributes interface{} `json:"decisionTaskStartedEventAttributes,omitempty"`
	Activitytaskcanceledeventattributes interface{} `json:"activityTaskCanceledEventAttributes,omitempty"`
	Lambdafunctionstartedeventattributes interface{} `json:"lambdaFunctionStartedEventAttributes,omitempty"`
	Timercanceledeventattributes interface{} `json:"timerCanceledEventAttributes,omitempty"`
	Timerstartedeventattributes interface{} `json:"timerStartedEventAttributes,omitempty"`
	Workflowexecutionterminatedeventattributes interface{} `json:"workflowExecutionTerminatedEventAttributes,omitempty"`
	Lambdafunctioncompletedeventattributes interface{} `json:"lambdaFunctionCompletedEventAttributes,omitempty"`
	Scheduleactivitytaskfailedeventattributes interface{} `json:"scheduleActivityTaskFailedEventAttributes,omitempty"`
	Activitytaskfailedeventattributes interface{} `json:"activityTaskFailedEventAttributes,omitempty"`
	Workflowexecutionfailedeventattributes interface{} `json:"workflowExecutionFailedEventAttributes,omitempty"`
	Requestcancelexternalworkflowexecutioninitiatedeventattributes interface{} `json:"requestCancelExternalWorkflowExecutionInitiatedEventAttributes,omitempty"`
	Externalworkflowexecutioncancelrequestedeventattributes interface{} `json:"externalWorkflowExecutionCancelRequestedEventAttributes,omitempty"`
	Starttimerfailedeventattributes interface{} `json:"startTimerFailedEventAttributes,omitempty"`
	Workflowexecutionsignaledeventattributes interface{} `json:"workflowExecutionSignaledEventAttributes,omitempty"`
	Canceltimerfailedeventattributes interface{} `json:"cancelTimerFailedEventAttributes,omitempty"`
	Childworkflowexecutiontimedouteventattributes interface{} `json:"childWorkflowExecutionTimedOutEventAttributes,omitempty"`
	Startlambdafunctionfailedeventattributes interface{} `json:"startLambdaFunctionFailedEventAttributes,omitempty"`
	Activitytaskstartedeventattributes interface{} `json:"activityTaskStartedEventAttributes,omitempty"`
	Failworkflowexecutionfailedeventattributes interface{} `json:"failWorkflowExecutionFailedEventAttributes,omitempty"`
	Lambdafunctionscheduledeventattributes interface{} `json:"lambdaFunctionScheduledEventAttributes,omitempty"`
	Workflowexecutioncompletedeventattributes interface{} `json:"workflowExecutionCompletedEventAttributes,omitempty"`
	Eventtype interface{} `json:"eventType"`
	Childworkflowexecutionfailedeventattributes interface{} `json:"childWorkflowExecutionFailedEventAttributes,omitempty"`
	Startchildworkflowexecutionfailedeventattributes interface{} `json:"startChildWorkflowExecutionFailedEventAttributes,omitempty"`
	Activitytaskscheduledeventattributes interface{} `json:"activityTaskScheduledEventAttributes,omitempty"`
	Lambdafunctiontimedouteventattributes interface{} `json:"lambdaFunctionTimedOutEventAttributes,omitempty"`
	Childworkflowexecutioncanceledeventattributes interface{} `json:"childWorkflowExecutionCanceledEventAttributes,omitempty"`
	Completeworkflowexecutionfailedeventattributes interface{} `json:"completeWorkflowExecutionFailedEventAttributes,omitempty"`
	Cancelworkflowexecutionfailedeventattributes interface{} `json:"cancelWorkflowExecutionFailedEventAttributes,omitempty"`
	Activitytaskcancelrequestedeventattributes interface{} `json:"activityTaskCancelRequestedEventAttributes,omitempty"`
	Lambdafunctionfailedeventattributes interface{} `json:"lambdaFunctionFailedEventAttributes,omitempty"`
	Decisiontaskscheduledeventattributes interface{} `json:"decisionTaskScheduledEventAttributes,omitempty"`
	Requestcancelactivitytaskfailedeventattributes interface{} `json:"requestCancelActivityTaskFailedEventAttributes,omitempty"`
	Externalworkflowexecutionsignaledeventattributes interface{} `json:"externalWorkflowExecutionSignaledEventAttributes,omitempty"`
	Childworkflowexecutioncompletedeventattributes interface{} `json:"childWorkflowExecutionCompletedEventAttributes,omitempty"`
	Childworkflowexecutionstartedeventattributes interface{} `json:"childWorkflowExecutionStartedEventAttributes,omitempty"`
	Signalexternalworkflowexecutionfailedeventattributes interface{} `json:"signalExternalWorkflowExecutionFailedEventAttributes,omitempty"`
	Markerrecordedeventattributes interface{} `json:"markerRecordedEventAttributes,omitempty"`
	Activitytasktimedouteventattributes interface{} `json:"activityTaskTimedOutEventAttributes,omitempty"`
	Activitytaskcompletedeventattributes interface{} `json:"activityTaskCompletedEventAttributes,omitempty"`
	Requestcancelexternalworkflowexecutionfailedeventattributes interface{} `json:"requestCancelExternalWorkflowExecutionFailedEventAttributes,omitempty"`
	Workflowexecutioncancelrequestedeventattributes interface{} `json:"workflowExecutionCancelRequestedEventAttributes,omitempty"`
}

// WorkflowExecutionCanceledEventAttributes represents the WorkflowExecutionCanceledEventAttributes schema from the OpenAPI specification
type WorkflowExecutionCanceledEventAttributes struct {
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Details interface{} `json:"details,omitempty"`
}

// DescribeWorkflowTypeInput represents the DescribeWorkflowTypeInput schema from the OpenAPI specification
type DescribeWorkflowTypeInput struct {
	Workflowtype interface{} `json:"workflowType"`
	Domain interface{} `json:"domain"`
}

// DeprecateDomainInput represents the DeprecateDomainInput schema from the OpenAPI specification
type DeprecateDomainInput struct {
	Name interface{} `json:"name"`
}

// FailWorkflowExecutionFailedEventAttributes represents the FailWorkflowExecutionFailedEventAttributes schema from the OpenAPI specification
type FailWorkflowExecutionFailedEventAttributes struct {
	Cause interface{} `json:"cause"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
}

// ListDomainsInput represents the ListDomainsInput schema from the OpenAPI specification
type ListDomainsInput struct {
	Registrationstatus interface{} `json:"registrationStatus"`
	Reverseorder interface{} `json:"reverseOrder,omitempty"`
	Maximumpagesize interface{} `json:"maximumPageSize,omitempty"`
	Nextpagetoken interface{} `json:"nextPageToken,omitempty"`
}

// StartTimerFailedEventAttributes represents the StartTimerFailedEventAttributes schema from the OpenAPI specification
type StartTimerFailedEventAttributes struct {
	Timerid interface{} `json:"timerId"`
	Cause interface{} `json:"cause"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
}

// WorkflowTypeDetail represents the WorkflowTypeDetail schema from the OpenAPI specification
type WorkflowTypeDetail struct {
	Configuration interface{} `json:"configuration"`
	Typeinfo interface{} `json:"typeInfo"`
}

// CompleteWorkflowExecutionFailedEventAttributes represents the CompleteWorkflowExecutionFailedEventAttributes schema from the OpenAPI specification
type CompleteWorkflowExecutionFailedEventAttributes struct {
	Cause interface{} `json:"cause"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
}

// FailWorkflowExecutionDecisionAttributes represents the FailWorkflowExecutionDecisionAttributes schema from the OpenAPI specification
type FailWorkflowExecutionDecisionAttributes struct {
	Details interface{} `json:"details,omitempty"`
	Reason interface{} `json:"reason,omitempty"`
}

// DecisionTaskTimedOutEventAttributes represents the DecisionTaskTimedOutEventAttributes schema from the OpenAPI specification
type DecisionTaskTimedOutEventAttributes struct {
	Scheduledeventid interface{} `json:"scheduledEventId"`
	Startedeventid interface{} `json:"startedEventId"`
	Timeouttype interface{} `json:"timeoutType"`
}

// SignalExternalWorkflowExecutionDecisionAttributes represents the SignalExternalWorkflowExecutionDecisionAttributes schema from the OpenAPI specification
type SignalExternalWorkflowExecutionDecisionAttributes struct {
	Signalname interface{} `json:"signalName"`
	Workflowid interface{} `json:"workflowId"`
	Control interface{} `json:"control,omitempty"`
	Input interface{} `json:"input,omitempty"`
	Runid interface{} `json:"runId,omitempty"`
}

// SignalWorkflowExecutionInput represents the SignalWorkflowExecutionInput schema from the OpenAPI specification
type SignalWorkflowExecutionInput struct {
	Runid interface{} `json:"runId,omitempty"`
	Signalname interface{} `json:"signalName"`
	Workflowid interface{} `json:"workflowId"`
	Domain interface{} `json:"domain"`
	Input interface{} `json:"input,omitempty"`
}

// ListOpenWorkflowExecutionsInput represents the ListOpenWorkflowExecutionsInput schema from the OpenAPI specification
type ListOpenWorkflowExecutionsInput struct {
	Domain interface{} `json:"domain"`
	Executionfilter interface{} `json:"executionFilter,omitempty"`
	Maximumpagesize interface{} `json:"maximumPageSize,omitempty"`
	Nextpagetoken interface{} `json:"nextPageToken,omitempty"`
	Reverseorder interface{} `json:"reverseOrder,omitempty"`
	Starttimefilter interface{} `json:"startTimeFilter"`
	Tagfilter interface{} `json:"tagFilter,omitempty"`
	Typefilter interface{} `json:"typeFilter,omitempty"`
}

// RequestCancelWorkflowExecutionInput represents the RequestCancelWorkflowExecutionInput schema from the OpenAPI specification
type RequestCancelWorkflowExecutionInput struct {
	Workflowid interface{} `json:"workflowId"`
	Domain interface{} `json:"domain"`
	Runid interface{} `json:"runId,omitempty"`
}

// RegisterDomainInput represents the RegisterDomainInput schema from the OpenAPI specification
type RegisterDomainInput struct {
	Workflowexecutionretentionperiodindays interface{} `json:"workflowExecutionRetentionPeriodInDays"`
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name"`
	Tags interface{} `json:"tags,omitempty"`
}

// ListActivityTypesInput represents the ListActivityTypesInput schema from the OpenAPI specification
type ListActivityTypesInput struct {
	Maximumpagesize interface{} `json:"maximumPageSize,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Nextpagetoken interface{} `json:"nextPageToken,omitempty"`
	Registrationstatus interface{} `json:"registrationStatus"`
	Reverseorder interface{} `json:"reverseOrder,omitempty"`
	Domain interface{} `json:"domain"`
}

// DecisionTaskCompletedEventAttributes represents the DecisionTaskCompletedEventAttributes schema from the OpenAPI specification
type DecisionTaskCompletedEventAttributes struct {
	Executioncontext interface{} `json:"executionContext,omitempty"`
	Scheduledeventid interface{} `json:"scheduledEventId"`
	Startedeventid interface{} `json:"startedEventId"`
}

// ScheduleActivityTaskFailedEventAttributes represents the ScheduleActivityTaskFailedEventAttributes schema from the OpenAPI specification
type ScheduleActivityTaskFailedEventAttributes struct {
	Activitytype interface{} `json:"activityType"`
	Cause interface{} `json:"cause"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Activityid interface{} `json:"activityId"`
}

// StartChildWorkflowExecutionDecisionAttributes represents the StartChildWorkflowExecutionDecisionAttributes schema from the OpenAPI specification
type StartChildWorkflowExecutionDecisionAttributes struct {
	Tasklist interface{} `json:"taskList,omitempty"`
	Input interface{} `json:"input,omitempty"`
	Childpolicy interface{} `json:"childPolicy,omitempty"`
	Executionstarttoclosetimeout interface{} `json:"executionStartToCloseTimeout,omitempty"`
	Lambdarole interface{} `json:"lambdaRole,omitempty"`
	Taskstarttoclosetimeout interface{} `json:"taskStartToCloseTimeout,omitempty"`
	Workflowtype interface{} `json:"workflowType"`
	Taskpriority interface{} `json:"taskPriority,omitempty"`
	Workflowid interface{} `json:"workflowId"`
	Control interface{} `json:"control,omitempty"`
	Taglist interface{} `json:"tagList,omitempty"`
}

// DomainConfiguration represents the DomainConfiguration schema from the OpenAPI specification
type DomainConfiguration struct {
	Workflowexecutionretentionperiodindays interface{} `json:"workflowExecutionRetentionPeriodInDays"`
}

// TaskList represents the TaskList schema from the OpenAPI specification
type TaskList struct {
	Name interface{} `json:"name"`
}

// TerminateWorkflowExecutionInput represents the TerminateWorkflowExecutionInput schema from the OpenAPI specification
type TerminateWorkflowExecutionInput struct {
	Details interface{} `json:"details,omitempty"`
	Domain interface{} `json:"domain"`
	Reason interface{} `json:"reason,omitempty"`
	Runid interface{} `json:"runId,omitempty"`
	Workflowid interface{} `json:"workflowId"`
	Childpolicy interface{} `json:"childPolicy,omitempty"`
}

// StartChildWorkflowExecutionFailedEventAttributes represents the StartChildWorkflowExecutionFailedEventAttributes schema from the OpenAPI specification
type StartChildWorkflowExecutionFailedEventAttributes struct {
	Initiatedeventid interface{} `json:"initiatedEventId"`
	Workflowid interface{} `json:"workflowId"`
	Workflowtype interface{} `json:"workflowType"`
	Cause interface{} `json:"cause"`
	Control interface{} `json:"control,omitempty"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
}

// RecordMarkerFailedEventAttributes represents the RecordMarkerFailedEventAttributes schema from the OpenAPI specification
type RecordMarkerFailedEventAttributes struct {
	Cause interface{} `json:"cause"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Markername interface{} `json:"markerName"`
}

// TagFilter represents the TagFilter schema from the OpenAPI specification
type TagFilter struct {
	Tag interface{} `json:"tag"`
}

// ActivityTaskStatus represents the ActivityTaskStatus schema from the OpenAPI specification
type ActivityTaskStatus struct {
	Cancelrequested interface{} `json:"cancelRequested"`
}

// LambdaFunctionCompletedEventAttributes represents the LambdaFunctionCompletedEventAttributes schema from the OpenAPI specification
type LambdaFunctionCompletedEventAttributes struct {
	Result interface{} `json:"result,omitempty"`
	Scheduledeventid interface{} `json:"scheduledEventId"`
	Startedeventid interface{} `json:"startedEventId"`
}

// WorkflowExecutionFailedEventAttributes represents the WorkflowExecutionFailedEventAttributes schema from the OpenAPI specification
type WorkflowExecutionFailedEventAttributes struct {
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Details interface{} `json:"details,omitempty"`
	Reason interface{} `json:"reason,omitempty"`
}

// CancelWorkflowExecutionDecisionAttributes represents the CancelWorkflowExecutionDecisionAttributes schema from the OpenAPI specification
type CancelWorkflowExecutionDecisionAttributes struct {
	Details interface{} `json:"details,omitempty"`
}

// StartTimerDecisionAttributes represents the StartTimerDecisionAttributes schema from the OpenAPI specification
type StartTimerDecisionAttributes struct {
	Control interface{} `json:"control,omitempty"`
	Starttofiretimeout interface{} `json:"startToFireTimeout"`
	Timerid interface{} `json:"timerId"`
}

// CompleteWorkflowExecutionDecisionAttributes represents the CompleteWorkflowExecutionDecisionAttributes schema from the OpenAPI specification
type CompleteWorkflowExecutionDecisionAttributes struct {
	Result interface{} `json:"result,omitempty"`
}

// RespondActivityTaskCompletedInput represents the RespondActivityTaskCompletedInput schema from the OpenAPI specification
type RespondActivityTaskCompletedInput struct {
	Tasktoken interface{} `json:"taskToken"`
	Result interface{} `json:"result,omitempty"`
}

// WorkflowExecutionCancelRequestedEventAttributes represents the WorkflowExecutionCancelRequestedEventAttributes schema from the OpenAPI specification
type WorkflowExecutionCancelRequestedEventAttributes struct {
	Externalworkflowexecution interface{} `json:"externalWorkflowExecution,omitempty"`
	Cause interface{} `json:"cause,omitempty"`
	Externalinitiatedeventid interface{} `json:"externalInitiatedEventId,omitempty"`
}

// ActivityTypeDetail represents the ActivityTypeDetail schema from the OpenAPI specification
type ActivityTypeDetail struct {
	Configuration interface{} `json:"configuration"`
	Typeinfo interface{} `json:"typeInfo"`
}

// ChildWorkflowExecutionCompletedEventAttributes represents the ChildWorkflowExecutionCompletedEventAttributes schema from the OpenAPI specification
type ChildWorkflowExecutionCompletedEventAttributes struct {
	Workflowtype interface{} `json:"workflowType"`
	Initiatedeventid interface{} `json:"initiatedEventId"`
	Result interface{} `json:"result,omitempty"`
	Startedeventid interface{} `json:"startedEventId"`
	Workflowexecution interface{} `json:"workflowExecution"`
}

// WorkflowExecutionTerminatedEventAttributes represents the WorkflowExecutionTerminatedEventAttributes schema from the OpenAPI specification
type WorkflowExecutionTerminatedEventAttributes struct {
	Cause interface{} `json:"cause,omitempty"`
	Childpolicy interface{} `json:"childPolicy"`
	Details interface{} `json:"details,omitempty"`
	Reason interface{} `json:"reason,omitempty"`
}

// WorkflowExecutionSignaledEventAttributes represents the WorkflowExecutionSignaledEventAttributes schema from the OpenAPI specification
type WorkflowExecutionSignaledEventAttributes struct {
	Externalinitiatedeventid interface{} `json:"externalInitiatedEventId,omitempty"`
	Externalworkflowexecution interface{} `json:"externalWorkflowExecution,omitempty"`
	Input interface{} `json:"input,omitempty"`
	Signalname interface{} `json:"signalName"`
}

// ActivityTaskScheduledEventAttributes represents the ActivityTaskScheduledEventAttributes schema from the OpenAPI specification
type ActivityTaskScheduledEventAttributes struct {
	Starttoclosetimeout interface{} `json:"startToCloseTimeout,omitempty"`
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Taskpriority interface{} `json:"taskPriority,omitempty"`
	Activityid interface{} `json:"activityId"`
	Scheduletoclosetimeout interface{} `json:"scheduleToCloseTimeout,omitempty"`
	Scheduletostarttimeout interface{} `json:"scheduleToStartTimeout,omitempty"`
	Heartbeattimeout interface{} `json:"heartbeatTimeout,omitempty"`
	Tasklist interface{} `json:"taskList"`
	Activitytype interface{} `json:"activityType"`
	Control interface{} `json:"control,omitempty"`
	Input interface{} `json:"input,omitempty"`
}

// RespondActivityTaskFailedInput represents the RespondActivityTaskFailedInput schema from the OpenAPI specification
type RespondActivityTaskFailedInput struct {
	Details interface{} `json:"details,omitempty"`
	Reason interface{} `json:"reason,omitempty"`
	Tasktoken interface{} `json:"taskToken"`
}

// ContinueAsNewWorkflowExecutionDecisionAttributes represents the ContinueAsNewWorkflowExecutionDecisionAttributes schema from the OpenAPI specification
type ContinueAsNewWorkflowExecutionDecisionAttributes struct {
	Input interface{} `json:"input,omitempty"`
	Lambdarole interface{} `json:"lambdaRole,omitempty"`
	Taglist interface{} `json:"tagList,omitempty"`
	Tasklist interface{} `json:"taskList,omitempty"`
	Workflowtypeversion interface{} `json:"workflowTypeVersion,omitempty"`
	Childpolicy interface{} `json:"childPolicy,omitempty"`
	Executionstarttoclosetimeout interface{} `json:"executionStartToCloseTimeout,omitempty"`
	Taskpriority interface{} `json:"taskPriority,omitempty"`
	Taskstarttoclosetimeout interface{} `json:"taskStartToCloseTimeout,omitempty"`
}

// RespondDecisionTaskCompletedInput represents the RespondDecisionTaskCompletedInput schema from the OpenAPI specification
type RespondDecisionTaskCompletedInput struct {
	Tasktoken interface{} `json:"taskToken"`
	Decisions interface{} `json:"decisions,omitempty"`
	Executioncontext interface{} `json:"executionContext,omitempty"`
}

// ChildWorkflowExecutionTimedOutEventAttributes represents the ChildWorkflowExecutionTimedOutEventAttributes schema from the OpenAPI specification
type ChildWorkflowExecutionTimedOutEventAttributes struct {
	Initiatedeventid interface{} `json:"initiatedEventId"`
	Startedeventid interface{} `json:"startedEventId"`
	Timeouttype interface{} `json:"timeoutType"`
	Workflowexecution interface{} `json:"workflowExecution"`
	Workflowtype interface{} `json:"workflowType"`
}

// DecisionTaskStartedEventAttributes represents the DecisionTaskStartedEventAttributes schema from the OpenAPI specification
type DecisionTaskStartedEventAttributes struct {
	Identity interface{} `json:"identity,omitempty"`
	Scheduledeventid interface{} `json:"scheduledEventId"`
}

// ExecutionTimeFilter represents the ExecutionTimeFilter schema from the OpenAPI specification
type ExecutionTimeFilter struct {
	Latestdate interface{} `json:"latestDate,omitempty"`
	Oldestdate interface{} `json:"oldestDate"`
}

// DeprecateActivityTypeInput represents the DeprecateActivityTypeInput schema from the OpenAPI specification
type DeprecateActivityTypeInput struct {
	Activitytype interface{} `json:"activityType"`
	Domain interface{} `json:"domain"`
}

// RegisterActivityTypeInput represents the RegisterActivityTypeInput schema from the OpenAPI specification
type RegisterActivityTypeInput struct {
	Defaulttaskstarttoclosetimeout interface{} `json:"defaultTaskStartToCloseTimeout,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Version interface{} `json:"version"`
	Defaulttaskheartbeattimeout interface{} `json:"defaultTaskHeartbeatTimeout,omitempty"`
	Domain interface{} `json:"domain"`
	Defaulttaskpriority interface{} `json:"defaultTaskPriority,omitempty"`
	Name interface{} `json:"name"`
	Defaulttasklist interface{} `json:"defaultTaskList,omitempty"`
	Defaulttaskscheduletoclosetimeout interface{} `json:"defaultTaskScheduleToCloseTimeout,omitempty"`
	Defaulttaskscheduletostarttimeout interface{} `json:"defaultTaskScheduleToStartTimeout,omitempty"`
}

// WorkflowExecutionCompletedEventAttributes represents the WorkflowExecutionCompletedEventAttributes schema from the OpenAPI specification
type WorkflowExecutionCompletedEventAttributes struct {
	Decisiontaskcompletedeventid interface{} `json:"decisionTaskCompletedEventId"`
	Result interface{} `json:"result,omitempty"`
}

// ChildWorkflowExecutionFailedEventAttributes represents the ChildWorkflowExecutionFailedEventAttributes schema from the OpenAPI specification
type ChildWorkflowExecutionFailedEventAttributes struct {
	Workflowexecution interface{} `json:"workflowExecution"`
	Workflowtype interface{} `json:"workflowType"`
	Details interface{} `json:"details,omitempty"`
	Initiatedeventid interface{} `json:"initiatedEventId"`
	Reason interface{} `json:"reason,omitempty"`
	Startedeventid interface{} `json:"startedEventId"`
}

// WorkflowExecutionInfo represents the WorkflowExecutionInfo schema from the OpenAPI specification
type WorkflowExecutionInfo struct {
	Cancelrequested interface{} `json:"cancelRequested,omitempty"`
	Workflowtype interface{} `json:"workflowType"`
	Closestatus interface{} `json:"closeStatus,omitempty"`
	Execution interface{} `json:"execution"`
	Executionstatus interface{} `json:"executionStatus"`
	Parent interface{} `json:"parent,omitempty"`
	Starttimestamp interface{} `json:"startTimestamp"`
	Taglist interface{} `json:"tagList,omitempty"`
	Closetimestamp interface{} `json:"closeTimestamp,omitempty"`
}

// ActivityTaskStartedEventAttributes represents the ActivityTaskStartedEventAttributes schema from the OpenAPI specification
type ActivityTaskStartedEventAttributes struct {
	Identity interface{} `json:"identity,omitempty"`
	Scheduledeventid interface{} `json:"scheduledEventId"`
}

// CountOpenWorkflowExecutionsInput represents the CountOpenWorkflowExecutionsInput schema from the OpenAPI specification
type CountOpenWorkflowExecutionsInput struct {
	Domain interface{} `json:"domain"`
	Executionfilter interface{} `json:"executionFilter,omitempty"`
	Starttimefilter interface{} `json:"startTimeFilter"`
	Tagfilter interface{} `json:"tagFilter,omitempty"`
	Typefilter interface{} `json:"typeFilter,omitempty"`
}
