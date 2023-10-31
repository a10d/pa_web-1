using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using TodolistApi;
using TodolistApi.Models;

namespace TodolistApi.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class TodoTypesController : ControllerBase
    {
        private readonly TodoContext _context;

        public TodoTypesController(TodoContext context)
        {
            _context = context;
        }

        // GET: api/TodoTypes
        [HttpGet]
        public async Task<ActionResult<IEnumerable<TodoType>>> GetTodoTypes()
        {
          if (_context.TodoTypes == null)
          {
              return NotFound();
          }
            return await _context.TodoTypes.ToListAsync();
        }

        // GET: api/TodoTypes/5
        [HttpGet("{id}")]
        public async Task<ActionResult<TodoType>> GetTodoType(int id)
        {
          if (_context.TodoTypes == null)
          {
              return NotFound();
          }
            var todoType = await _context.TodoTypes.FindAsync(id);

            if (todoType == null)
            {
                return NotFound();
            }

            return todoType;
        }

        // PUT: api/TodoTypes/5
        // To protect from overposting attacks, see https://go.microsoft.com/fwlink/?linkid=2123754
        [HttpPut("{id}")]
        public async Task<IActionResult> PutTodoType(int id, TodoType todoType)
        {
            if (id != todoType.Id)
            {
                return BadRequest();
            }

            _context.Entry(todoType).State = EntityState.Modified;

            try
            {
                await _context.SaveChangesAsync();
            }
            catch (DbUpdateConcurrencyException)
            {
                if (!TodoTypeExists(id))
                {
                    return NotFound();
                }
                else
                {
                    throw;
                }
            }

            return NoContent();
        }

        // POST: api/TodoTypes
        // To protect from overposting attacks, see https://go.microsoft.com/fwlink/?linkid=2123754
        [HttpPost]
        public async Task<ActionResult<TodoType>> PostTodoType(TodoType todoType)
        {
          if (_context.TodoTypes == null)
          {
              return Problem("Entity set 'TodoContext.TodoTypes'  is null.");
          }
            _context.TodoTypes.Add(todoType);
            await _context.SaveChangesAsync();

            return CreatedAtAction("GetTodoType", new { id = todoType.Id }, todoType);
        }

        // DELETE: api/TodoTypes/5
        [HttpDelete("{id}")]
        public async Task<IActionResult> DeleteTodoType(int id)
        {
            if (_context.TodoTypes == null)
            {
                return NotFound();
            }
            var todoType = await _context.TodoTypes.FindAsync(id);
            if (todoType == null)
            {
                return NotFound();
            }

            _context.TodoTypes.Remove(todoType);
            await _context.SaveChangesAsync();

            return NoContent();
        }

        private bool TodoTypeExists(int id)
        {
            return (_context.TodoTypes?.Any(e => e.Id == id)).GetValueOrDefault();
        }
    }
}
