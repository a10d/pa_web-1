<?php

namespace App\Http\Controllers\Task;

use App\Foundation\Http\Controller;
use App\Models\Task;
use Illuminate\Http\Response;

class DeleteTask extends Controller
{
    public function __invoke(Task $task)
    {
        $this->authorize('delete', $task);

        $task->delete();

        return response(null, Response::HTTP_ACCEPTED);
    }
}
