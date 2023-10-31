<?php

namespace App\Http\Controllers\TaskType;

use App\Http\Requests\TaskType\CreateTaskType as CreateTaskTypeRequest;
use App\Models\TaskType;
use Illuminate\Http\JsonResponse;

class CreateTaskType
{
    public function __invoke(CreateTaskTypeRequest $request): JsonResponse
    {
        $taskType = TaskType::create($request->validated());

        return response()->json([
            'message' => 'Task created successfully',
            'taskType' => $taskType,
        ]);
    }

}
