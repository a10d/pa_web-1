<?php

namespace App\Http\Controllers\Task;

use App\Foundation\Http\Controller;
use App\Models\Task;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Response;
use Throwable;

class DeleteTask extends Controller
{
    /**
     * @throws Throwable
     */
    public function __invoke(Task $task): JsonResponse
    {
        $this->authorize('delete', $task);

        $task->deleteOrFail();

        return response()->json([
            'message' => 'Task deleted successfully',
        ]);
    }
}
