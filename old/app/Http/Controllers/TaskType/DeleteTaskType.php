<?php

namespace App\Http\Controllers\TaskType;

use App\Foundation\Http\Controller;
use App\Models\TaskType;
use Illuminate\Http\JsonResponse;
use Throwable;

class DeleteTaskType extends Controller
{
    /**
     * @throws Throwable
     */
    public function __invoke(TaskType $taskType): JsonResponse
    {
        $this->authorize('delete', [$taskType]);

        $taskType->deleteOrFail();

        return response()->json([
            'message' => 'Task type deleted successfully',
        ]);
    }

}
