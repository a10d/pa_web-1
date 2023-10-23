<?php

namespace App\Http\Controllers\Task;

use App\Foundation\Http\Controller;
use App\Http\Requests\Task\CreateTask as CreateTaskRequest;
use App\Models\Task;
use Illuminate\Http\JsonResponse;
use Illuminate\Support\Arr;
use Throwable;

class CreateTask extends Controller
{
    /**
     * @throws Throwable
     */
    public function __invoke(CreateTaskRequest $request): JsonResponse
    {
        $task = new Task();

        $task
            ->fill(Arr::except($request->validated(), ['assignees']))
            ->saveOrFail();

        $task->assignees = $request->get('assignees');

        return response()->json([
            'message' => 'Task created successfully',
            'task' => $task,
        ]);
    }
}
