<?php

namespace App\Http\Controllers\TaskType;

use app\Http\Requests\TaskType\UpdateTaskType as UpdateTaskTypeRequest;
use App\Models\TaskType;
use Illuminate\Http\JsonResponse;
use Throwable;

class UpdateTaskType
{
    /**
     * @throws Throwable
     */
    public function __invoke(TaskType $taskType, UpdateTaskTypeRequest $request): JsonResponse
    {
        $taskType->fill($request->validated());

        $taskType->saveOrFail();

        return response()->json([
            'message' => 'Task type updated successfully',
            'taskType' => $taskType,
        ]);
    }
}
