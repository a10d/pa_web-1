<?php

namespace App\Http\Controllers\Task;

use App\Foundation\Http\Controller;
use App\Http\Requests\Task\UpdateTask as UpdateTaskRequest;
use App\Models\Task;
use Illuminate\Http\JsonResponse;
use Throwable;

class UpdateTask extends Controller
{
    /**
     * @throws Throwable
     */
    public function __invoke(Task $task, UpdateTaskRequest $request): JsonResponse
    {
        $task->fill($request->validated());

        $task->saveOrFail();

        return response()->json([
            'message' => 'Task updated successfully',
            'task' => $task,
        ]);
    }
}
