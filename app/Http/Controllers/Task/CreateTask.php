<?php

namespace App\Http\Controllers\Task;

use App\Foundation\Http\Controller;
use App\Http\Requests\Task\CreateTask as CreateTaskRequest;
use App\Models\Task;
use Illuminate\Http\JsonResponse;

class CreateTask extends Controller
{
    public function __invoke(CreateTaskRequest $request): JsonResponse
    {
        $task = Task::create($request->validated());

        return response()->json([
            'message' => 'Task created successfully',
            'task' => $task,
        ]);
    }
}
